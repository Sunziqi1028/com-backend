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
	ID       uint64    `gorm:"column:id"`
	Name     string    `gorm:"column:name"`
	Code     int       `gorm:"column:code"`
	Category int       `gorm:"column:category"`
	CreateAt time.Time `gorm:"column:create_at;autoCreateTime"`
	UpdateAt time.Time `gorm:"column:update_at;autoUpdateTime:milli"`
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
