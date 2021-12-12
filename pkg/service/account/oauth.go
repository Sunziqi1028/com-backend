package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	"ceres/pkg/utility/auth"
	"ceres/pkg/utility/jwt"
	"errors"

	"gorm.io/gorm"

	"github.com/gotomicro/ego/core/elog"
)

// LoginWithOauth common oauth login logic in comunion
func LoginWithOauth(client auth.OauthClient, oauthType account.ComerAccountType) (response *account.ComerLoginResponse, err error) {
	oauth, err := client.GetUserProfile()
	if err != nil {
		return
	}
	// try to find account
	comerAccount, err := account.GetComerAccount(mysql.DB, oauthType, oauth.GetUserID())
	if err != nil {
		elog.Errorf("Comunion Oauth login faild, because of %v", err)
		return
	}
	//set default profile status
	var isProfiled bool
	var comerProfile account.ComerProfile
	var comer account.Comer

	if comerAccount.ID == 0 {
		comer = account.Comer{}
		err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
			if er = account.CreateComer(tx, &comer); er != nil {
				return er
			}
			comerAccount = account.ComerAccount{
				ComerID:   comer.ID,
				OIN:       oauth.GetUserID(),
				IsPrimary: true,
				Nick:      oauth.GetUserNick(),
				Avatar:    oauth.GetUserAvatar(),
				Type:      oauthType,
				IsLinked:  true,
			}
			if er = account.CreateAccount(tx, &comerAccount); er != nil {
				return er
			}
			return
		})
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

	// sign with jwt using the comer UIN
	token := jwt.Sign(comer.ID)

	address := ""
	if comer.Address != nil {
		address = *comer.Address
	}

	response = &account.ComerLoginResponse{
		IsProfiled: isProfiled,
		Avatar:     "",
		Name:       comerProfile.Name,
		Address:    address,
		Token:      token,
	}

	return
}

// LinkOauthAccountToComer link a new Oauth account to the current comer
func LinkOauthAccountToComer(ComerID uint64, client auth.OauthClient, oauthType account.ComerAccountType) (err error) {
	oauth, err := client.GetUserProfile()
	comerAccount, err := account.GetComerAccount(mysql.DB, oauthType, oauth.GetUserID())
	if err != nil {
		elog.Errorf("Comunion Oauth login faild, because of %v", err)
		return
	}
	if comerAccount.ID != 0 {
		err = errors.New("account has bind to another comer")
		elog.Error(err.Error())
		return err
	}
	comerAccount = account.ComerAccount{
		ComerID:   ComerID,
		OIN:       oauth.GetUserID(),
		IsPrimary: true,
		Nick:      oauth.GetUserNick(),
		Avatar:    oauth.GetUserAvatar(),
		Type:      oauthType,
		IsLinked:  true,
	}
	if err = account.CreateAccount(mysql.DB, &comerAccount); err != nil {
		return err
	}
	return
}
