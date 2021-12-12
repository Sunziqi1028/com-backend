package account

import (
	"ceres/pkg/model"
)

type ComerAccountType int

const (
	GithubOauth   ComerAccountType = 1
	TwitterOauth  ComerAccountType = 2
	FacebookOauth ComerAccountType = 3
	LikedinOauth  ComerAccountType = 4
	GoogleOauth   ComerAccountType = 5
)

// Comer the comer model of comunion inner account
type Comer struct {
	model.Base
	Address *string `gorm:"column:address"`
}

// TableName Comer table name for gorm
func (Comer) TableName() string {
	return "comer"
}

// ComerAccount the account model of comer
type ComerAccount struct {
	model.Base
	ComerID   uint64           `gorm:"column:comer_id"`
	OIN       string           `gorm:"column:oin"`
	IsPrimary bool             `gorm:"column:is_primary"`
	Nick      string           `gorm:"column:nick"`
	Avatar    string           `gorm:"column:avatar"`
	Type      ComerAccountType `gorm:"column:type"`
	IsLinked  bool             `gorm:"column:is_linked"`
}

// TableName the ComerAccount table name for gorm
func (ComerAccount) TableName() string {
	return "comer_account"
}

// ComerProfile model
type ComerProfile struct {
	model.Base
	ComerID  uint64 `gorm:"column:comer_id" json:"comerID"`
	Name     string `gorm:"column:name" json:"name"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Location string `gorm:"column:location" json:"location"`
	Website  string `gorm:"column:website" json:"website"`
	BIO      string `gorm:"column:bio" json:"bio"`
}

// TableName the Profile table name for gorm
func (ComerProfile) TableName() string {
	return "comer_profile"
}

// Skill model
type Skill struct {
	model.Base
	Name string `gorm:"column:name"`
}

// TableName the Skill table name for gorm
func (Skill) TableName() string {
	return "comer_skill"
}

// ComerSkillRel model
type ComerSkillRel struct {
	model.Base
	ComerID uint64 `gorm:"column:comer_id"`
	SkillID uint64 `gorm:"column:skill_id"`
}

// TableName the comer skill relation table name for gorm
func (ComerSkillRel) TableName() string {
	return "comer_skill_rel"
}
