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
	"time"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/qiniu/x/log"
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
			return
		}
		log.Warn(err)
		return err
	}
	//verify wallet and nonce
	if err = VerifyEthWallet(address, nonce, signature); err != nil {
		return
	}
	var comer account.Comer
	if err = account.GetComerByAddress(mysql.DB, address, &comer); err != nil {
		log.Warn(err)
		return err
	}
	//set default profile status
	var isProfiled bool
	var profile account.ComerProfile

	if comer.ID == 0 {
		comer = account.Comer{
			Address: &address,
		}
		// create a new comer
		err = account.CreateComer(mysql.DB, &comer)
		if err != nil {
			return
		}
		isProfiled = false
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
	}

	return
}

// LinkEthAccountToComer link a new eth wallet account to comer
func LinkEthAccountToComer(comerID uint64, address, signature string) (err error) {
	nonce, err := redis.Client.Get(context.TODO(), address)
	if err != nil {
		if err.Error() == "eredis get string error eredis exec command get fail, redis: nil" {
			err = router.ErrBadRequest.WithMsg("Please get nonce")
			return
		}
		log.Warn(err)
		return err
	}
	//verify wallet and nonce
	if err = VerifyEthWallet(address, nonce, signature); err != nil {
		return
	}

	var comer account.Comer
	if err = account.GetComerByID(mysql.DB, comerID, &comer); err != nil {
		return
	}
	if comer.Address != nil {
		return router.ErrBadRequest.WithMsg("Current comer has linked with a wallet")
	}

	if err = account.GetComerByAddress(mysql.DB, address, &comer); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
	}
	if comer.ID != 0 {
		return router.ErrBadRequest.WithMsg("Current eth wallet account is linked with a comer")
	}

	if err = account.UpdateComerAddress(mysql.DB, comerID, address); err != nil {
		log.Warn(err)
		return
	}

	_, err = redis.Client.Del(context.TODO(), address)
	if err != nil {
		log.Warnf("redis remove nonce key failed %v", err)
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
