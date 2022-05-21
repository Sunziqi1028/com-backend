package startup

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/startup"

	"github.com/qiniu/x/log"
)

func FollowStartup(ComerID, startupID uint64) (err error) {
	return model.CreateStartupFollowRel(mysql.DB, ComerID, startupID)
}

func UnfollowStartup(ComerID, startupID uint64) (err error) {
	followRel := model.FollowRelation{
		StartupID: startupID,
		ComerID:   ComerID,
	}
	if err = model.DeleteStartupFollowRel(mysql.DB, &followRel); err != nil {
		log.Warn(err)
	}
	return
}

func ListFollowStartups(ComerID uint64, request *model.ListStartupRequest, response *model.ListStartupsResponse) (err error) {
	var startups []model.Startup
	total, err := model.ListFollowedStartups(mysql.DB, ComerID, request, &startups)
	if err != nil {
		log.Warn(err)
		return
	}
	response.Total = total
	response.List = startups
	return
}
