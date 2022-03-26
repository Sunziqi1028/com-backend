package startup

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/startup"

	"github.com/qiniu/x/log"
)

// ListStartups get current comer accounts
func ListStartups(comerID uint64, request *model.ListStartupRequest, response *model.ListStartupsResponse) (err error) {
	var startups []model.Startup
	total, err := model.ListStartups(mysql.DB, comerID, request, &startups)
	if err != nil {
		log.Warn(err)
		return
	}
	if total == 0 {
		response.List = make([]model.Startup, 0)
		return
	}
	response.List = startups
	response.Total = total
	return
}

func GetStartup(startupID uint64, response *model.GetStartupResponse) (err error) {
	var startup model.Startup
	if err = model.GetStartup(mysql.DB, startupID, &response.Startup); err != nil {
		log.Warn(err)
		return
	}
	response.Startup = startup
	return
}

func StartupNameIsExist(name string) (isExist bool, err error) {
	isExist, err = model.StartupNameIsExist(mysql.DB, name)
	if err != nil {
		log.Warn(err)
		return
	}
	return
}

func StartupTokenContractIsExist(tokenContract string) (isExist bool, err error) {
	isExist, err = model.StartupTokenContractIsExist(mysql.DB, tokenContract)
	if err != nil {
		log.Warn(err)
		return
	}
	return
}
