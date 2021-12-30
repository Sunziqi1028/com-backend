package account

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/account"

	"github.com/qiniu/x/log"
)

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
