package account

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/utility/auth"
	"ceres/pkg/utility/jwt"

	"github.com/qiniu/x/log"
	"gorm.io/gorm"
)

// LoginWithOauth common oauth login logic in comunion
func LoginWithOauth(client auth.OauthClient, oauthType account.ComerAccountType, response *account.ComerLoginResponse) (err error) {
	oauth, err := client.GetUserProfile()
	if err != nil {
		log.Warn(err)
		return router.ErrBadRequest.WithMsg("Login authorization failed")
	}

	// try to find account
	var comerAccount account.ComerAccount
	if err = account.GetComerAccount(mysql.DB, oauthType, oauth.GetUserID(), &comerAccount); err != nil {
		log.Warn(err)
		return
	}

	//set default profile status
	var isProfiled bool
	var profile account.ComerProfile
	var comer account.Comer

	if comerAccount.ID == 0 {
		comer = account.Comer{}
		err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
			if er = account.CreateComer(tx, &comer); er != nil {
				log.Warn(er)
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
				log.Warn(er)
				return er
			}
			return
		})
		if err != nil {
			return
		}
	} else {
		if err = account.GetComerProfile(mysql.DB, comer.ID, &profile); err != nil {
			log.Warn(err)
			return err
		}
		if profile.ID != 0 {
			isProfiled = true
		}
	}

	// sign with jwt using the comer UIN
	token := jwt.Sign(comer.ID)

	address := ""
	if comer.Address != nil {
		address = *comer.Address
	}

	*response = account.ComerLoginResponse{
		IsProfiled: isProfiled,
		Avatar:     profile.Avatar,
		Name:       profile.Name,
		Address:    address,
		Token:      token,
	}

	return nil
}

// LinkOauthAccountToComer link a new Oauth account to the current comer
func LinkOauthAccountToComer(ComerID uint64, client auth.OauthClient, oauthType account.ComerAccountType) (err error) {
	oauth, err := client.GetUserProfile()
	if err != nil {
		log.Warn(err)
		return
	}
	// try to find account
	var comerAccount account.ComerAccount
	if err = account.GetComerAccount(mysql.DB, oauthType, oauth.GetUserID(), &comerAccount); err != nil {
		log.Warn(err)
		return err
	}
	if comerAccount.ID != 0 {
		err = router.ErrBadRequest.WithMsg("Account has bind to another comer")
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
