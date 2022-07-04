package model

import (
	"ceres/pkg/initialization/utility"
	"math"
	"strings"
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
	ID        uint64    `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updatedAt"`
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

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" || strings.TrimSpace(p.Sort) == "" {
		p.Sort = "created_at desc"
	}
	return p.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
