package event

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	model "ceres/pkg/model/startup"
	"ceres/pkg/model/tag"
	"errors"

	"github.com/ethereum/go-ethereum/common"

	"gorm.io/gorm"

	"github.com/qiniu/x/log"
)

func HandleStartup(address string, startupProto interface{}) {
	startupTemp := startupProto.(struct {
		Name          string         `json:"name"`
		Mode          uint8          `json:"mode"`
		Hashtag       []string       `json:"hashtag"`
		Logo          string         `json:"logo"`
		Mission       string         `json:"mission"`
		TokenContract common.Address `json:"tokenContract"`
		Wallets       []struct {
			Name          string         `json:"name"`
			WalletAddress common.Address `json:"walletAddress"`
		} `json:"wallets"`
		Overview   string `json:"overview"`
		IsValidate bool   `json:"isValidate"`
	})

	comer := account.Comer{}
	if err := account.GetComerByAddress(mysql.DB, address, &comer); err != nil {
		log.Warn(err)
		return
	}
	if comer.ID == 0 {
		log.Warn(errors.New("comer does not exit"))
		return
	}
	startup := model.Startup{
		ComerID:              comer.ID,
		Name:                 startupTemp.Name,
		Mode:                 model.Mode(startupTemp.Mode),
		Logo:                 startupTemp.Logo,
		Mission:              startupTemp.Mission,
		TokenContractAddress: startupTemp.TokenContract.String(),
		Overview:             startupTemp.Overview,
	}
	if err := mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		//create startup
		if er = model.CreateStartup(tx, &startup); er != nil {
			return
		}
		var wallets []model.Wallet
		for _, v := range startupTemp.Wallets {
			wallets = append(wallets, model.Wallet{
				ComerID:       comer.ID,
				StartupID:     startup.ID,
				WalletName:    v.Name,
				WalletAddress: v.WalletAddress.String(),
			})
		}
		//create startup wallet
		if er = model.CreateStartupWallet(tx, wallets); er != nil {
			return
		}
		//create skill
		var tagRelList []tag.TagTargetRel
		for _, skillName := range startupTemp.Hashtag {
			var isIndex bool
			if len(skillName) > 2 && skillName[0:1] == "#" {
				isIndex = true
			}
			skill := tag.Tag{
				Name:     skillName,
				IsIndex:  isIndex,
				Category: tag.Startup,
			}
			if er = tag.FirstOrCreateTag(tx, &skill); er != nil {
				return er
			}
			tagRelList = append(tagRelList, tag.TagTargetRel{
				TagID:    skill.ID,
				Target:   tag.Startup,
				TargetID: startup.ID,
			})
		}
		//batch create comer skill relation
		if len(tagRelList) == 0 {
			return er
		}
		if er = tag.BatchCreateTagRel(tx, tagRelList); er != nil {
			log.Warn(er)
			return er
		}
		return er
	}); err != nil {
		log.Warn(err)
		return
	}
	return
}
