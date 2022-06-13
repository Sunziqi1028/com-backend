package account

import (
	"ceres/pkg/model"
	"ceres/pkg/model/tag"
)

type ComerAccountType int

const (
	GithubOauth   ComerAccountType = 1
	GoogleOauth   ComerAccountType = 2
	TwitterOauth  ComerAccountType = 3
	FacebookOauth ComerAccountType = 4
	LikedinOauth  ComerAccountType = 5
)

// Comer the comer model of comunion inner account
type Comer struct {
	model.Base
	Address *string `gorm:"column:address" json:"address"`
}

// TableName Comer table name for gorm
func (Comer) TableName() string {
	return "comer"
}

// ComerAccount the account model of comer
type ComerAccount struct {
	model.Base
	ComerID   uint64           `gorm:"column:comer_id" json:"comerID"`
	OIN       string           `gorm:"column:oin" json:"oin"`
	IsPrimary bool             `gorm:"column:is_primary" json:"isPrimary"`
	Nick      string           `gorm:"column:nick" json:"nick"`
	Avatar    string           `gorm:"column:avatar" json:"avatar"`
	Type      ComerAccountType `gorm:"column:type" json:"type"`
	IsLinked  bool             `gorm:"column:is_linked" json:"isLinked"`
}

// TableName the ComerAccount table name for gorm
func (ComerAccount) TableName() string {
	return "comer_account"
}

// ComerProfile model
type ComerProfile struct {
	model.Base
	ComerID  uint64    `gorm:"column:comer_id" json:"comerID"`
	Name     string    `gorm:"column:name" json:"name"`
	Avatar   string    `gorm:"column:avatar" json:"avatar"`
	Location string    `gorm:"column:location" json:"location"`
	TimeZone string    `gorm:"column:time_zone" json:"timeZone"`
	Website  string    `gorm:"column:website" json:"website"`
	Email    string    `gorm:"column:email" json:"email"`
	Twitter  string    `gorm:"column:twitter" json:"twitter"`
	Discord  string    `gorm:"column:discord" json:"discord"`
	Telegram string    `gorm:"column:telegram" json:"telegram"`
	Medium   string    `gorm:"column:medium" json:"medium"`
	BIO      string    `gorm:"column:bio" json:"bio"`
	Skills   []tag.Tag `gorm:"many2many:tag_target_rel;foreignKey:ComerID;joinForeignKey:TargetID;" json:"skills"`
}

// TableName the Profile table name for gorm
func (ComerProfile) TableName() string {
	return "comer_profile"
}
