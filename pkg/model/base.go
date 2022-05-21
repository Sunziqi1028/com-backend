package model

import (
	"ceres/pkg/initialization/utility"
	"time"

	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uint64    `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDeleted bool      `gorm:"column:is_deleted" json:"isDeleted"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = utility.Sequence.Next()
	return
}

// RelationBase contains common columns for all tables.
type RelationBase struct {
	ID        uint64 `gorm:"primary_key;column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (base *RelationBase) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = utility.Sequence.Next()
	return
}

// ListRequest list request
type ListRequest struct {
	Limit     int  `form:"limit" binding:"gt=0"`
	Offset    int  `form:"offset" binding:"gte=0"`
	IsDeleted bool `form:"isDeleted"`
}
