package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"

	"github.com/qiniu/x/log"
)

// UserInfo  unlink the comer account
func UserInfo(comerID uint64, UserInfo *model.ComerLoginResponse) error {
	var comer model.Comer
	if err := model.GetComerByID(mysql.DB, comerID, &comer); err != nil {
		log.Warn(err)
		return err
	}
	if comer.Address != nil {
		UserInfo.Address = *comer.Address
	} else {
		UserInfo.Address = ""
	}

	var isProfiled bool
	var profile model.ComerProfile
	if err := model.GetComerProfile(mysql.DB, comerID, &profile); err != nil {
		log.Warn(err)
		return err
	}

	if profile.ID != 0 {
		isProfiled = true
	}

	UserInfo.Nick = profile.Name
	UserInfo.Avatar = profile.Avatar
	UserInfo.IsProfiled = isProfiled
	UserInfo.ComerID = comer.ID

	return nil
}

// GetComerAccounts get current comer accounts
func GetComerAccounts(comerID uint64, response *model.ComerOuterAccountListResponse) (err error) {
	accountList := make([]model.ComerAccount, 0)
	if err = model.ListAccount(mysql.DB, comerID, &accountList); err != nil {
		log.Warn(err)
		return
	}
	*response = model.ComerOuterAccountListResponse{
		List:  accountList,
		Total: uint64(len(accountList)),
	}
	return
}

// UnlinkComerAccount  unlink the comer account
func UnlinkComerAccount(comerID, accountID uint64) error {
	return model.DeleteAccount(mysql.DB, comerID, accountID)
}

// GetComerInfo get comer info
func GetComerInfo(comerID uint64, response *model.GetComerInfoResponse) (err error) {
	if err = model.GetComerProfile(mysql.DB, comerID, &response.ComerProfile); err != nil {
		log.Warn(err)
		return err
	}
	if err = model.GetComerByID(mysql.DB, comerID, &response.Comer); err != nil {
		log.Warn(err)
		return err
	}
	return
}
