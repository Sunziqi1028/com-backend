package meta

import (
	"ceres/pkg/model/account"
	"time"

	"github.com/jinzhu/gorm"
)

/// Profile the comer profile model
type ProfileSkill struct {
	ID       uint64    `gorm:"id"`
	UIN      uint64    `gorm:"uin"`
	Name     string    `gorm:"name"`
	IsValid  bool      `gorm:"valid"`
	CreateAt time.Time `gorm:"create_at"`
	UpdateAt time.Time `gorm:"update_at"`
}

func (ProfileSkill) TableName() string {
	return "comer_profile_skill_tag_tbl"
}

/// CreateSkillTag
/// create a new skill tag
func CreateSkillTag(db *gorm.DB, skill *ProfileSkill) {
	db.Save(skill)
}

/// ListAllSkillTags
/// list all skill tags
func ListAllSkillTags(db *gorm.DB) (skills []ProfileSkill, err error) {
	r := db.Find(&skills)
	err = r.Error
	return
}

/// GetSkills
/// return the skills list to current profile
func GetSkills(db *gorm.DB, profile *account.Profile) (skills []ProfileSkill) {
	// skillsID := strings.Split(profile.Skills, ",")
	// TODO: should find more effective for this method implementation
	return
}
