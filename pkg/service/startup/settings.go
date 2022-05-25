package startup

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/startup"
	"ceres/pkg/model/tag"
	"ceres/pkg/router"
	"errors"
	"gorm.io/gorm"
	"time"

	"github.com/qiniu/x/log"
)

// UpdateStartupBasicSetting update startup security and social setting
func UpdateStartupBasicSetting(startupID uint64, request *model.UpdateStartupBasicSettingRequest) (err error) {
	//get startup
	var startup model.Startup
	if err = model.GetStartup(mysql.DB, startupID, &startup); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if startup.ID == 0 {
		return router.ErrBadRequest.WithMsg("startup does not exist")
	}
	var tagIds []uint64
	var tagRelList []tag.TagTargetRel
	startupBasicSetting := model.BasicSetting{
		KYC:           *request.KYC,
		ContractAudit: *request.ContractAudit,
		Website:       *request.Website,
		Discord:       *request.Discord,
		Twitter:       *request.Twitter,
		Telegram:      *request.Telegram,
		Docs:          *request.Docs,
	}
	if err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		for _, tagName := range request.HashTags {
			var isIndex bool
			if len(tagName) > 2 && tagName[0:1] == "#" {
				isIndex = true
			}
			hashTag := tag.Tag{
				Name:     tagName,
				Category: tag.Startup,
				IsIndex:  isIndex,
			}
			if er = tag.FirstOrCreateTag(tx, &hashTag); er != nil {
				return er
			}
			tagRelList = append(tagRelList, tag.TagTargetRel{
				TagID:    hashTag.ID,
				Target:   tag.Startup,
				TargetID: startupID,
			})
			tagIds = append(tagIds, hashTag.ID)
		}
		//delete not used hashtags
		if er = tag.DeleteTagRel(tx, startupID, tag.Startup, tagIds); er != nil {
			return er
		}
		//batch create startup hashtag rel
		if er = tag.BatchCreateTagRel(tx, tagRelList); er != nil {
			return er
		}
		//update startup basic setting
		if er = model.UpdateStartupBasicSetting(tx, startupID, &startupBasicSetting); er != nil {
			return er
		}
		return er
	}); err != nil {
		log.Warn(err)
		return
	}
	return
}

func UpdateStartupFinanceSetting(startupID, comerID uint64, request *model.UpdateStartupFinanceSettingRequest) (err error) {
	//get startup
	var startup model.Startup
	if err = model.GetStartup(mysql.DB, startupID, &startup); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if startup.ID == 0 {
		return router.ErrBadRequest.WithMsg("startup does not exist")
	}
	var walletIds []uint64
	var walletList []model.Wallet
	startupFinanceSetting := model.FinanceSetting{
		TokenContractAddress: *request.TokenContractAddress,
		LaunchNetwork:        *request.LaunchNetwork,
		TokenName:            *request.TokenName,
		TokenSymbol:          *request.TokenSymbol,
		TotalSupply:          *request.TotalSupply,
		PresaleStart:         ConverToDatetime(*request.PresaleStart),
		PresaleEnd:           ConverToDatetime(*request.PresaleEnd),
		LaunchDate:           ConverToDatetime(*request.LaunchDate),
	}

	if err = mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		for _, v := range request.Wallets {
			wallet := model.Wallet{
				ComerID:       comerID,
				StartupID:     startup.ID,
				WalletName:    v.WalletName,
				WalletAddress: v.WalletAddress,
			}
			if er = model.FirstOrCreateWallet(tx, &wallet); er != nil {
				return er
			}
			wallet.WalletName = v.WalletName
			wallet.WalletAddress = v.WalletAddress
			walletList = append(walletList, wallet)
			walletIds = append(walletIds, wallet.ID)
		}
		//batch update startup wallet
		if er = model.BatchUpdateStartupWallet(tx, walletList); er != nil {
			return er
		}
		//delete not used startup wallet
		if er = model.DeleteStartupWallet(tx, startupID, walletIds); er != nil {
			return er
		}
		//update startup finance setting
		if er = model.UpdateStartupFinanceSetting(tx, startupID, &startupFinanceSetting); er != nil {
			return er
		}
		return er
	}); err != nil {
		log.Warn(err)
		return
	}
	return
}

func ConverToDatetime(strTime string) (t time.Time) {
	var err error
	const timeFormat = "2006-01-02T15:04:05Z"
	if t, err = time.ParseInLocation(timeFormat, strTime, time.UTC); err != nil {
		if t, err = time.ParseInLocation("2006-01-02", strTime, time.UTC); err != nil {
			t = time.Time{}
			return
		}
	}
	return t
}
