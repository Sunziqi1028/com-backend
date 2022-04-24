package startup

import (
	"ceres/pkg/model"
	"ceres/pkg/model/tag"
)

type Mode uint8

const (
	ModeESG Mode = 1
	ModeNGO Mode = 2
	ModeDAO Mode = 3
	ModeCOM Mode = 4
)

type Startup struct {
	model.Base
	ComerID              uint64    `gorm:"comer_id" json:"comerID"`
	Name                 string    `gorm:"name" json:"name"`
	Mode                 Mode      `gorm:"mode" json:"mode"`
	Logo                 string    `gorm:"logo" json:"logo"`
	Mission              string    `gorm:"mission" json:"mission"`
	TokenContractAddress string    `gorm:"token_contract_address" json:"tokenContractAddress"`
	Overview             string    `gorm:"overview" json:"overview"`
	IsSet                bool      `gorm:"is_set" json:"isSet"`
	KYC                  string    `gorm:"kyc" json:"kyc"`
	ContractAudit        string    `gorm:"contract_audit" json:"contractAudit"`
	HashTags             []tag.Tag `gorm:"many2many:tag_target_rel;foreignKey:ID;joinForeignKey:TargetID;" json:"hashTags"`
	Website              string    `gorm:"website" json:"website"`
	Discord              string    `gorm:"discord" json:"discord"`
	Twitter              string    `gorm:"twitter" json:"twitter"`
	Telegram             string    `gorm:"telegram" json:"telegram"`
	Docs                 string    `gorm:"docs" json:"docs"`
	PresaleDate          string    `gorm:"presale_date" json:"presaleDate"`
	LaunchDate           string    `gorm:"launch_date" json:"launchDate"`
	Wallets              []Wallet  `json:"wallets"`
}

// TableName Startup table name for gorm
func (Startup) TableName() string {
	return "startup"
}

type Wallet struct {
	model.Base
	ComerID       uint64 `gorm:"comer_id" json:"comerID"`
	StartupID     uint64 `gorm:"startup_id" json:"startupID"`
	WalletName    string `gorm:"wallet_name" json:"walletName"`
	WalletAddress string `gorm:"wallet_address" json:"walletAddress"`
}

// TableName wallet table name for gorm
func (Wallet) TableName() string {
	return "startup_wallet"
}

type FollowRelation struct {
	model.RelationBase
	ComerID   uint64 `gorm:"comer_id" json:"comerID"`
	StartupID uint64 `gorm:"startup_id" json:"startupID"`
}

// TableName Followed table name for gorm
func (FollowRelation) TableName() string {
	return "startup_follow_rel"
}

// Startup security and social setting
type StartupBasicSetting struct {
	KYC           string `gorm:"kyc" json:"kyc"`
	ContractAudit string `gorm:"contract_audit" json:"contractAudit"`
	Website       string `gorm:"website" json:"website"`
	Discord       string `gorm:"discord" json:"discord"`
	Twitter       string `gorm:"twitter" json:"twitter"`
	Telegram      string `gorm:"telegram" json:"telegram"`
	Docs          string `gorm:"docs" json:"docs"`
}

// Startup finance setting
type StartupFinanceSetting struct {
	TokenContractAddress string `gorm:"token_contract_address" json:"tokenContractAddress"`
	PresaleDate          string `gorm:"presale_date" json:"presaleDate"`
	LaunchDate           string `gorm:"launch_date" json:"launchDate"`
	//Wallets              []Wallet  `json:"wallets"`
}
