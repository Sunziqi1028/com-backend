package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/initialization/redis"
	"ceres/pkg/initialization/utility"
	"ceres/pkg/model/account"
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
	"github.com/gotomicro/ego/core/elog"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
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

// VerifyEthLogin verify the signature and login with the eth wallet
// FIXME: have to check the address nonce in redis to keep the request is vaild 
func VerifyEthLogin(address, messageHash, signature string, walletType int) (response *account.ComerLoginResponse, err error) {
	addrKey := common.HexToAddress(address)
	sig := hexutil.MustDecode(signature)
	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}
	// if not end with 27 then will be error in message in the response
	hash := hexutil.MustDecode(messageHash)

	pubKey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return
	}

	recoverAddr := crypto.PubkeyToAddress(*pubKey)

	if recoverAddr != addrKey {
		err = errors.New("Not match the origin public key")
		return
	}

	comer, err := account.GetComerByAccountOIN(mysql.DB, address)
	if err != nil {
		elog.Error(err.Error())
		return
	}

	if comer.ID == 0 {
		// create a new comer with the origin ID
		// create comer with account
		now := time.Now()
		comer.UIN = utility.AccountSequnece.Next()
		comer.ComerID = strings.Replace(uuid.Must(uuid.NewV4(), nil).String(), "-", "", -1)
		comer.Avatar = comer.ComerID
		comer.Nick = address
		comer.CreateAt = now
		comer.UpdateAt = now
		outer := &account.Account{}
		outer.Identifier = utility.AccountSequnece.Next()
		outer.OIN = address
		outer.UIN = comer.UIN
		outer.IsMain = true
		outer.IsLinked = true
		outer.Nick = comer.Nick
		outer.Avatar = comer.Avatar
		outer.Category = account.EthAccount
		outer.Type = walletType
		outer.CreateAt = now
		outer.UpdateAt = now
		// Create the account and comer within transaction
		err = account.CreateComerWithAccount(mysql.DB, &comer, outer)
		if err != nil {
			elog.Errorf("Comunion Eth login faild, because of %v", err)
			return
		}
		_, err = redis.Client.Del(context.TODO(), address)
		if err != nil {
			elog.Errorf("Comunion redis remove key failed %v", err)
		}
	}

	// sign with jwt
	token := jwt.Sign(comer.UIN)

	response = &account.ComerLoginResponse{
		ComerID: comer.ComerID,
		Address: comer.Address,
		Nick:    comer.Nick,
		Avatar:  comer.Avatar,
		Token:   token,
		UIN:     comer.UIN,
	}
	return
}

// LinkEthAccountToComer link a new eth wallet account to comer
// FIXME: have to check the address nonce in redis to keep the request is vaild 
func LinkEthAccountToComer(uin uint64, address, messageHash, signature string, walletType int) (err error) {
	err = mysql.DB.Transaction(func(tx *gorm.DB) (err error) {
		comer, err := account.GetComerByAccountUIN(tx, uin)
		if err != nil {
			return
		}
		if comer.ID == 0 {
			return errors.New("comer is not exists")
		}
		addrKey := common.HexToAddress(address)
		sig := hexutil.MustDecode(signature)
		if sig[64] == 27 || sig[64] == 28 {
			sig[64] -= 27
		}
		// if not end with 27 then will be error in message in the response
		hash := hexutil.MustDecode(messageHash)

		pubKey, err := crypto.SigToPub(hash, sig)
		if err != nil {
			return
		}

		recoverAddr := crypto.PubkeyToAddress(*pubKey)

		if recoverAddr != addrKey {
			err = errors.New("Not match the origin public key")
			return
		}
		// origin address passed from the frontend is the 0x prefix
		// but in this logic ceres will move the 0x predix in the router to do next
		refComer, err := account.GetComerByAccountOIN(mysql.DB, address)
		if err != nil {
			elog.Error(err.Error())
			return err
		}

		if refComer.ID == 0 {
			outer, err := account.GetAccountByOIN(tx, address)
			if err != nil {
				return err
			}
			if outer.ID == 0 {
				outer.Identifier = utility.AccountSequnece.Next()
			}
			now := time.Now()
			outer.OIN = address
			outer.UIN = comer.UIN
			outer.IsMain = false
			outer.IsLinked = true
			outer.Nick = comer.Nick
			outer.Avatar = comer.Avatar
			outer.Category = account.EthAccount
			outer.Type = walletType
			outer.CreateAt = now
			outer.UpdateAt = now
			// Create the account and comer within transaction
			err = account.LinkComerWithAccount(mysql.DB, uin, &outer)
			if err != nil {
				return err
			}
			_, err = redis.Client.Del(context.TODO(), address)
			if err != nil {
				elog.Errorf("Comunion redis remove key failed %v", err)
				return nil
			}
			return nil
		}
		return errors.New("current eth wallet account is linked with a comer")
	})
	return
}
