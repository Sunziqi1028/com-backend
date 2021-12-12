package model

import (
	"ceres/pkg/initialization/utility"
	"time"

	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uint64 `gorm:"primary_key;column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool `gorm:"column:is_deleted"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = utility.Sequence.Next()
	return
}
