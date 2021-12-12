package account

import (
	"ceres/pkg/model/account"
	"ceres/pkg/utility/auth"
)

// LoginWithOauth common oauth login logic in comunion
func LoginWithOauth(client auth.OauthClient, oauthType account.ComerAccountType) (response *account.ComerLoginResponse, err error) {
	//oauth, err := client.GetUserProfile()
	//if err != nil {
	//	return
	//}
	//// try to find account
	//comerAccount, err := account.GetComerByAccountOIN(mysql.DB, oauthType, oauth.GetUserID())
	//if err != nil {
	//	elog.Errorf("Comunion Oauth login faild, because of %v", err)
	//	return
	//}
	//var comer account.Comer
	//
	//if comerAccount.ID == 0 {
	//	now := time.Now()
	//	comer = account.Comer{
	//		ID:        utility.ComerSequnece.Next(),
	//		CreatedAt: now,
	//		UpdatedAt: now,
	//	}
	//	comerAccount = account.ComerAccount{
	//		ID:        utility.AccountSequnece.Next(),
	//		ComerID:   utility.ComerSequnece.Next(),
	//		OIN:       oauth.GetUserID(),
	//		IsPrimary: true,
	//		Nick:      oauth.GetUserNick(),
	//		Avatar:    oauth.GetUserAvatar(),
	//		Category:  account.OauthAccount,
	//		Type:      oauthType,
	//		IsLinked:  true,
	//		CreatedAt: now,
	//		UpdatedAt: now,
	//	}
	//	// Create the account and comer within transaction
	//	err = account.CreateComerWithAccount(mysql.DB, &comer, &comerAccount)
	//	if err != nil {
	//		elog.Errorf("Comunion Oauth login faild, because of %v", err)
	//		return
	//	}
	//}
	//
	//// sign with jwt using the comer UIN
	//token := jwt.Sign(comer.ID)
	//
	//response = &account.ComerLoginResponse{
	//	Address: comer.Address,
	//	Token:   token,
	//}

	return
}

// LinkOauthAccountToComer link a new Oauth account to the current comer
func LinkOauthAccountToComer(uin uint64, client auth.OauthClient, oauthType int) (err error) {
	//err = mysql.DB.Transaction(func(tx *gorm.DB) error {
	//	comer, err := account.GetComerByAccountUIN(tx, uin)
	//	if err != nil {
	//		return err
	//	}
	//	if comer.ID == 0 {
	//		return errors.New("comer is not exists")
	//	}
	//	oauth, err := client.GetUserProfile()
	//	if err != nil {
	//		return err
	//	}
	//	refComer, err := account.GetComerByAccountOIN(mysql.DB, oauth.GetUserID())
	//	if err != nil {
	//		return err
	//	}
	//	if refComer.ID == 0 {
	//		outer, err := account.GetAccountByOIN(tx, oauth.GetUserID())
	//		if err != nil {
	//			return err
	//		}
	//		if outer.ID == 0 {
	//			// if current account is not exists then create now
	//			outer.Identifier = utility.AccountSequnece.Next()
	//		}
	//		now := time.Now()
	//		outer.OIN = oauth.GetUserID()
	//		outer.UIN = comer.UIN
	//		outer.IsMain = false
	//		outer.IsLinked = true
	//		outer.Nick = comer.Nick
	//		outer.Avatar = comer.Avatar
	//		outer.Category = account.OauthAccount
	//		outer.Type = oauthType
	//		outer.CreateAt = now
	//		outer.UpdateAt = now
	//		err = account.LinkComerWithAccount(mysql.DB, uin, &outer)
	//		if err != nil {
	//			return err
	//		}
	//		return nil
	//	}
	//	return errors.New("current oauth account is linked with a comer")
	//})
	return
}
