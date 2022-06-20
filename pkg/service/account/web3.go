package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/initialization/redis"
	"ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/utility/jwt"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/qiniu/x/log"
	"gorm.io/gorm"
)

const (
	expire = time.Second * 240
)

// createNonce
// create a new nonce
func createNonce() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06v", rand.Intn(1000000000))
}

// GenerateWeb3LoginNonce
// generate the login nonce help frontend to sign the login signature
func GenerateWeb3LoginNonce(address string, response *account.WalletNonceResponse) (err error) {
	nonce, err := redis.Client.Get(context.TODO(), address)
	if err != nil {
		if err.Error() != "eredis get string error eredis exec command get fail, redis: nil" {
			log.Warn(err)
			return err
		}
	}
	if nonce == "" {
		nonce = createNonce()
		err = redis.Client.Set(context.TODO(), address, nonce, expire)
		if err != nil {
			log.Errorf("NONCE set fail %s", err)
		}
	}

	response.Nonce = nonce
	return
}

// LoginWithEthWallet common eth wallet login
func LoginWithEthWallet(address, signature string, response *account.ComerLoginResponse) (err error) {
	nonce, err := redis.Client.Get(context.TODO(), address)
	if err != nil {
		if err.Error() == "eredis get string error eredis exec command get fail, redis: nil" {
			err = router.ErrBadRequest.WithMsg("Please get nonce")
			return err
		}
		log.Warn(err)
		return err
	}
	//verify wallet and nonce
	if err = VerifyEthWallet(address, nonce, signature); err != nil {
		return err
	}
	var comer account.Comer
	if err = account.GetComerByAddress(mysql.DB, address, &comer); err != nil {
		log.Warn(err)
		return err
	}
	//set default profile status
	var (
		isProfiled bool
		profile    account.ComerProfile
		firstLogin = false
	)

	if comer.ID == 0 {
		comer = account.Comer{
			Address: &address,
		}
		// create a new comer
		err = account.CreateComer(mysql.DB, &comer)
		if err != nil {
			return err
		}
		isProfiled = false
		firstLogin = true
	} else {
		//get comer profile
		if err = account.GetComerProfile(mysql.DB, comer.ID, &profile); err != nil {
			log.Warn(err)
			return err
		}
		if profile.ID != 0 {
			isProfiled = true
		}
	}

	_, err = redis.Client.Del(context.TODO(), address)
	if err != nil {
		log.Warnf("Comunion redis remove key failed %v", err)
	}

	// sign with jwt
	token := jwt.Sign(comer.ID)

	*response = account.ComerLoginResponse{
		IsProfiled: isProfiled,
		Avatar:     profile.Avatar,
		Nick:       profile.Name,
		Address:    address,
		Token:      token,
		ComerID:    comer.ID,
		FirstLogin: firstLogin,
	}

	return
}

// LinkEthAccountToComer link a new eth wallet account to comer
func LinkEthAccountToComer(comerID uint64, address, signature string) (err error, finalComerId uint64) {
	finalComerId = comerID
	nonce, err := redis.Client.Get(context.TODO(), address)
	if err != nil {
		if err.Error() == "eredis get string error eredis exec command get fail, redis: nil" {
			log.Warn("Please get nonce")
			err = router.ErrBadRequest.WithMsg("Please get nonce")
			return
		}
		log.Warn(err)
		return
	}
	//verify wallet and nonce
	if err = VerifyEthWallet(address, nonce, signature); err != nil {
		log.Warn(err)
		return
	}

	var (
		targetComer         account.Comer
		targetComerAccounts []account.ComerAccount
		targetComerProfile  account.ComerProfile
		// targetComer 是否 注册完成(即有comerProfile)
		targetComerRegistedComplete = false
	)
	if err = account.GetComerByID(mysql.DB, comerID, &targetComer); err != nil {
		log.Warn(err)
		return
	}
	add := targetComer.Address
	if add != nil && strings.TrimSpace(*add) != "" {
		if strings.TrimSpace(*add) != address {
			log.Warn("Current targetComer has linked with a wallet")
			return router.ErrBadRequest.WithMsg("Current targetComer has linked with a wallet"), finalComerId
		}
		return
	}

	if err = account.GetComerAccountsByComerId(mysql.DB, comerID, &targetComerAccounts); err != nil {
		return
	}

	if err = account.GetComerProfile(mysql.DB, comerID, &targetComerProfile); err != nil {
		log.Warn(err)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn(err)
			return
		}
		targetComerRegistedComplete = false
	}
	if targetComerProfile.ID != 0 && !targetComerProfile.IsDeleted {
		targetComerRegistedComplete = true
	}

	var (
		comerByAddress                       account.Comer
		accountsOfComerByAddress             []account.ComerAccount
		comerByAddressHasSmeTypeOauthAccount bool
	)
	if err = account.GetComerByAddress(mysql.DB, address, &comerByAddress); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn(err)
			return
		}
	}
	if comerByAddress.ID != 0 {
		log.Info("Current eth wallet account is linked with another targetComer")
		if targetComerRegistedComplete {
			return router.ErrBadRequest.WithMsg("Current eth wallet account is linked with another targetComer"), finalComerId
		}

		if err = account.GetComerAccountsByComerId(mysql.DB, comerByAddress.ID, &accountsOfComerByAddress); err != nil {
			log.Warn(err)
			return
		}
		if accountsOfComerByAddress != nil && len(accountsOfComerByAddress) > 0 {
			for _, byAddress := range accountsOfComerByAddress {
				for _, comerAccount := range targetComerAccounts {
					if byAddress.Type == comerAccount.Type {
						comerByAddressHasSmeTypeOauthAccount = true
						break
					}
				}
			}
			if comerByAddressHasSmeTypeOauthAccount {
				return router.ErrInternalServer.WithMsg("Current eth wallet account is linked with another targetComer"), finalComerId
			}
		}
	}

	if targetComerRegistedComplete {
		if err = account.UpdateComerAddress(mysql.DB, comerID, address); err != nil {
			log.Warn(err)
			return
		}
	} else {
		for _, comerAccount := range targetComerAccounts {
			if comerByAddress.ID != 0 {
				if err = account.BindComerAccountToComerId(mysql.DB, comerAccount.ID, comerByAddress.ID); err != nil {
					log.Warn(err)
					return
				}
			} else {
				if err = account.UpdateComerAddress(mysql.DB, comerID, address); err != nil {
					log.Warn(err)
					return
				}
			}
		}
	}

	_, err = redis.Client.Del(context.TODO(), address)
	if err != nil {
		log.Warnf("redis remove nonce key failed %v\n", err)
	}
	return
}

// VerifyEthWallet verify the signature and login with the wallet
func VerifyEthWallet(address, nonce, signature string) error {
	addrKey := common.HexToAddress(address)
	sig := hexutil.MustDecode(signature)
	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(nonce), nonce)
	msg256 := crypto.Keccak256([]byte(msg))
	pubKey, err := crypto.SigToPub(msg256, sig)
	if err != nil {
		return err
	}
	recoverAddr := crypto.PubkeyToAddress(*pubKey)
	if recoverAddr != addrKey {
		err = router.ErrBadRequest.WithMsg("Address mismatch")
		return err
	}
	return nil
}
