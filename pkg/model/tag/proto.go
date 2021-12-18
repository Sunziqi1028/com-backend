package tag

import (
	"ceres/pkg/model"
)

type Target string

// const tag category type
const (
	ComerSkillTag Target = "comerSkill"
	StartupTag    Target = "startup"
)

// Tag  Comunion tag for startup bounty profile and other position need Tag.
type Tag struct {
	model.Base
	Name    string `gorm:"column:name"`
	IsIndex bool   `gorm:"column:is_index"`
}

// TableName identify the table name of this model.
func (Tag) TableName() string {
	return "tag"
}

// TagTargetRel  Comunion tag for startup bounty profile and other position need TagTargetRel.
type TagTargetRel struct {
	model.RelationBase
	TargetID uint64 `column:"target_id"`
	Target   Target `column:"target"`
	TagID    uint64 `column:"tag_id"`
}

// TableName identify the table name of this model.
func (TagTargetRel) TableName() string {
	return "tag_target_rel"
}
