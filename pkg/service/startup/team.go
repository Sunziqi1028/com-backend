package startup

import (
	"ceres/pkg/initialization/mysql"
	model "ceres/pkg/model/startup_team"

	"github.com/qiniu/x/log"
)

// ListStartupTeamMembers get startup team members comer
func ListStartupTeamMembers(startupID uint64, request *model.ListStartupTeamMemberRequest, response *model.ListStartupTeamMemberResponse) (err error) {
	total, err := model.ListStartupTeamMembers(mysql.DB, startupID, request, &response.List)
	if err != nil {
		log.Warn(err)
		return
	}
	if total == 0 {
		response.List = make([]model.StartupTeamMember, 0)
		return
	}
	response.Total = total
	return
}

// CreateStartupTeamMember create startup team member
func CreateStartupTeamMember(startupID, comerID uint64, request *model.CreateStartupTeamMemberRequest) (err error) {
	startupTeam := model.StartupTeamMember{
		StartupID: startupID,
		ComerID:   comerID,
		Position:  request.Position,
	}
	if err = model.CreateStartupTeamMembers(mysql.DB, &startupTeam); err != nil {
		log.Warn(err)
	}
	return
}

// UpdateStartupTeamMember update startup team member
func UpdateStartupTeamMember(startupID, comerID uint64, request *model.UpdateStartupTeamMemberRequest) (err error) {
	startupTeam := model.StartupTeamMember{
		StartupID: startupID,
		ComerID:   comerID,
		Position:  request.Position,
	}
	if err = model.UpdateStartupTeamMember(mysql.DB, &startupTeam); err != nil {
		log.Warn(err)
	}
	return
}

// DeleteStartupTeamMember delete startup team member
func DeleteStartupTeamMember(startupID, comerID uint64) (err error) {
	startupTeam := model.StartupTeamMember{
		StartupID: startupID,
		ComerID:   comerID,
	}
	if err = model.DeleteStartupTeamMember(mysql.DB, &startupTeam); err != nil {
		log.Warn(err)
	}
	return
}
