package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"
)

// manage the account informations link and unlink

// GetComerAccounts get current comer accounts
func GetComerAccounts(comerID uint64, response *model.ComerOuterAccountListResponse) (err error) {
	accountList := make([]model.ComerAccount, 0)
	if err = model.ListAccount(mysql.DB, comerID, &accountList); err != nil {
		return
	}
	*response = model.ComerOuterAccountListResponse{
		List:  accountList,
		Total: uint64(len(accountList)),
	}
	return
}

// UnlinkComerAccount  unlink the comer account
func UnlinkComerAccount(comerID, accountID uint64) (err error) {
	return model.DeleteAccount(mysql.DB, comerID, accountID)
}
