package event

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	model "ceres/pkg/model/startup"
	team "ceres/pkg/model/startup_team"
	"errors"

	"gorm.io/gorm"

	"github.com/qiniu/x/log"
)

func HandleStartup(address string, startupProto interface{}, txHash string) {
	log.Info("listen startup data: ", startupProto)
	startupTemp := startupProto.(struct {
		Name       string `json:"name"`
		Mode       uint8  `json:"mode"`
		Logo       string `json:"logo"`
		Mission    string `json:"mission"`
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
		ComerID:       comer.ID,
		Name:          startupTemp.Name,
		Mode:          model.Mode(startupTemp.Mode),
		Logo:          startupTemp.Logo,
		Mission:       startupTemp.Mission,
		Overview:      startupTemp.Overview,
		ContractAudit: txHash,
	}
	if err := mysql.DB.Transaction(func(tx *gorm.DB) (er error) {
		//create startup
		if er = model.CreateStartup(tx, &startup); er != nil {
			return
		}

		//create default team member
		teamMember := team.StartupTeamMember{
			StartupID: startup.ID,
			ComerID:   comer.ID,
			Position:  "founder",
		}
		if er = team.CreateStartupTeamMembers(mysql.DB, &teamMember); er != nil {
			return
		}

		return er
	}); err != nil {
		log.Warn(err)
		return
	}
	return
}
