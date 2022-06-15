package startup

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/startup"

	"github.com/qiniu/x/log"
)

// ListStartups get current comer accounts
func ListStartups(comerID uint64, request *model.ListStartupRequest, response *model.ListStartupsResponse) (err error) {
	total, err := model.ListStartups(mysql.DB, comerID, request, &response.List)
	if err != nil {
		log.Warn(err)
		return
	}
	if total == 0 {
		response.List = make([]model.Startup, 0)
		return
	}
	response.Total = total
	for i, startup := range response.List {
		response.List[i].MemberCount = len(startup.Members)
		response.List[i].FollowCount = len(startup.Follows)
	}
	return
}

func GetStartup(startupID uint64, response *model.GetStartupResponse) (err error) {
	if err = model.GetStartup(mysql.DB, startupID, &response.Startup); err != nil {
		log.Warn(err)
		return
	}
	response.MemberCount = len(response.Members)
	response.FollowCount = len(response.Follows)
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

func StartupFollowedByComer(startupID, comerID uint64) (isFollowed bool, err error) {
	isFollowed, err = model.StartupFollowIsExist(mysql.DB, startupID, comerID)
	if err != nil {
		log.Warn(err)
		return
	}
	return
}
