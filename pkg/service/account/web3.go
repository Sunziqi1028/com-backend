package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/initialization/redis"
	"ceres/pkg/model/account"
	"ceres/pkg/utility/jwt"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/gotomicro/ego/core/elog"
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
func GenerateWeb3LoginNonce(address string) (response *account.WalletNonceResponse, err error) {
	nonce, err := redis.Client.Get(context.TODO(), address)
	if err != nil {
		elog.Debugf("NONCE test is %s", err.Error())
	}

	if nonce == "" {
		nonce = createNonce()
		err = redis.Client.Set(context.TODO(), address, nonce, expire)
		if err != nil {
			elog.Warnf("NONCE set fail %s", err.Error())
		}
	}

	response = &account.WalletNonceResponse{Nonce: nonce}

	return
}

// LoginWithEthWallet common eth wallet login
func LoginWithEthWallet(address, signature, nonce string) (response *account.ComerLoginResponse, err error) {
	//verify wallet and nonce
	if err = VerifyEthWallet(address, nonce, signature); err != nil {
		return
	}

	comer, err := account.GetComerByAddress(mysql.DB, address)
	if err != nil {
		elog.Error(err.Error())
		return
	}
	//set default profile status
	var isProfiled bool
	var comerProfile account.ComerProfile

	if comer.ID == 0 {
		comer = account.Comer{
			Address: address,
		}
		// create a new comer
		err = account.CreateComer(mysql.DB, &comer)
		if err != nil {
			elog.Errorf("Comunion Eth login faild, because of %v", err)
			return
		}
		isProfiled = false
	} else {
		comerProfile, err = account.GetComerProfile(mysql.DB, comer.ID)
		if err != nil {
			elog.Errorf("Comunion get comer profile fauld, because of %v", err)
			return
		}
		if comerProfile.ID != 0 {
			isProfiled = true
		}
	}

	_, err = redis.Client.Del(context.TODO(), address)
	if err != nil {
		elog.Errorf("Comunion redis remove key failed %v", err)
	}

	// sign with jwt
	token := jwt.Sign(comer.ID)

	response = &account.ComerLoginResponse{
		Address:    comer.Address,
		Token:      token,
		Name:       comerProfile.Name,
		Avatar:     comerProfile.Avatar,
		IsProfiled: isProfiled,
	}
	return
}

// LinkEthAccountToComer link a new eth wallet account to comer
func LinkEthAccountToComer(comerID uint64, address, signature, nonce string) (err error) {
	//verify wallet and nonce
	if err = VerifyEthWallet(address, nonce, signature); err != nil {
		return
	}

	refComer, err := account.GetComerByID(mysql.DB, comerID)
	if err != nil {
		elog.Error(err.Error())
		return err
	}

	if refComer.Address != "" {
		return errors.New("Current comer has linked with a wallet")
	}

	refComer, err = account.GetComerByAddress(mysql.DB, address)
	if err != nil {
		elog.Error(err.Error())
		return err
	}

	if refComer.Address != "" {
		return errors.New("Current eth wallet account is linked with a comer")
	}

	if err = account.UpdateComerAddress(mysql.DB, comerID, address); err != nil {
		return
	}

	_, err = redis.Client.Del(context.TODO(), address)
	if err != nil {
		elog.Errorf("Comunion redis remove key failed %v", err)
	}
	return nil
}

// VerifyEthWallet verify the signature and login with the wallet
func VerifyEthWallet(address, nonce, signature string) (err error) {
	//addrKey := common.HexToAddress(address)
	//sig := hexutil.MustDecode(signature)
	//if sig[64] == 27 || sig[64] == 28 {
	//	sig[64] -= 27
	//}
	//msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(nonce), nonce)
	//msg256 := crypto.Keccak256([]byte(msg))
	//pubKey, err := crypto.SigToPub(msg256, sig)
	//if err != nil {
	//	return
	//}
	//recoverAddr := crypto.PubkeyToAddress(*pubKey)
	//if recoverAddr != addrKey {
	//	err = errors.New("Not match the origin public key")
	//	return
	//}
	return
}
