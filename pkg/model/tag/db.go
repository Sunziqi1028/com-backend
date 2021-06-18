package tag

import (
	"time"

	"github.com/jinzhu/gorm"
)

// const tag category type
const (
	StartupTag = 1
	SkillTag   = 2
)

// Tag  Comunion tag for startup bounty profile and other position need Tag.
type Tag struct {
	ID       uint64    `gorm:"id"`
	Name     string    `gorm:"name"`
	Code     int       `gorm:"code"`
	Category int       `gorm:"category"`
	CreateAt time.Time `gorm:"create_at"`
	UpdateAt time.Time `gorm:"update_at"`
}

// TableName identify the table name of this model.
func (Tag) TableName() string {
	return "comunion_tags_tbl"
}

// GetTagListByCategory  get some category tag list by the category code
// @see ceres/pkg/mode/tag/db.go const such as StartupTag SkillTag.
func GetTagListByCategory(db *gorm.DB, category int) (tags []Tag, err error) {
	db = db.Where("category = ?", category).Find(tags)
	err = db.Error
	if err != nil {
		return
	}
	return
}
