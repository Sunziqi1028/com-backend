package startup_team

import (
	"gorm.io/gorm"
)

// ListStartupTeamMembers  get startup team members
func ListStartupTeamMembers(db *gorm.DB, startupID uint64, input *ListStartupTeamMemberRequest, output *[]StartupTeamMember) (total int64, err error) {
	if startupID != 0 {
		db = db.Where("startup_id = ?", startupID)
	}
	if err = db.Table("startup_team_member_rel").Count(&total).Error; err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = db.Order("created_at DESC").Limit(input.Limit).Offset(input.Offset).Preload("Comer").Preload("ComerProfile").Find(output).Error
	return
}

// CreateStartupTeamMembers  add startup team members
func CreateStartupTeamMembers(db *gorm.DB, input *StartupTeamMember) (err error) {
	return db.Create(input).Error
}

// UpdateStartupTeamMember  update startup team member title
func UpdateStartupTeamMember(db *gorm.DB, input *StartupTeamMember) (err error) {
	return db.Table("startup_team_member_rel").Where("comer_id = ? AND startup_id = ?", input.ComerID, input.StartupID).Update("position", input.Position).Error
}

// DeleteStartupTeamMember delete startup team member
func DeleteStartupTeamMember(db *gorm.DB, input *StartupTeamMember) (err error) {
	return db.Where("comer_id = ? AND startup_id = ?", input.ComerID, input.StartupID).Delete(input).Error
}
