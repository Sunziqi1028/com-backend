package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/initialization/utility"
	"ceres/pkg/model/account"
	"ceres/pkg/utility/auth"
	"ceres/pkg/utility/jwt"
	"errors"

	"github.com/gotomicro/ego/core/elog"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// LoginWithOauth common oauth login logic in comunion
func LoginWithOauth(client auth.OauthClient, oauthType int) (response *account.ComerLoginResponse, err error) {
	oauth, err := client.GetUserProfile()
	if err != nil {
		return
	}
	// try to find comer
	comer, err := account.GetComerByAccountOIN(mysql.DB, oauth.GetUserID())
	if err != nil {
		elog.Errorf("Comunion Oauth login faild, because of %v", err)
		return
	}

	if comer.ID == 0 {
		// create comer with account
		comer.UIN = utility.AccountSequnece.Next()
		comer.Avatar = oauth.GetUserAvatar()
		comer.Nick = oauth.GetUserNick()
		comer.ComerID = uuid.Must(uuid.NewV4(), nil).String()
		if comer.Avatar == "" {
			comer.Avatar = comer.ComerID
		}

		outer := &account.Account{}

		outer.OIN = oauth.GetUserID()
		outer.UIN = comer.UIN
		outer.IsMain = true
		outer.IsLinked = true
		outer.Nick = comer.Nick
		outer.Avatar = comer.Avatar
		outer.Category = account.OauthAccount
		outer.Type = oauthType
		// Create the account and comer within transaction
		err = account.CreateComerWithAccount(mysql.DB, &comer, outer)
		if err != nil {
			elog.Errorf("Comunion Oauth login faild, because of %v", err)
			return
		}
	}

	// sign with jwt using the comer UIN

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

// LinkOauthAccountToComer link a new Oauth account to the current comer
func LinkOauthAccountToComer(uin uint64, client auth.OauthClient, oauthType int) (err error) {
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		comer, err := account.GetComerByAccountUIN(tx, uin)
		if err != nil {
			return err
		}
		if comer.ID == 0 {
			return errors.New("comer is not exists")
		}
		oauth, err := client.GetUserProfile()
		if err != nil {
			return err
		}
		refComer, err := account.GetComerByAccountOIN(mysql.DB, oauth.GetUserID())
		if err != nil {
			return err
		}
		if refComer.ID == 0 {
			outer, err := account.GetAccountByOIN(tx, oauth.GetUserID())
			if err != nil {
				return err
			}
			if outer.ID == 0 {
				// if current account is not exists then create now
				outer.Identifier = utility.AccountSequnece.Next()
			}
			outer.OIN = oauth.GetUserID()
			outer.UIN = comer.UIN
			outer.IsMain = false
			outer.IsLinked = true
			outer.Nick = comer.Nick
			outer.Avatar = comer.Avatar
			outer.Category = account.OauthAccount
			outer.Type = oauthType
			err = account.LinkComerWithAccount(mysql.DB, uin, &outer)
			if err != nil {
				return err
			}
			return nil
		}
		return errors.New("current oauth account is linked with a comer")
	})
	return
}
