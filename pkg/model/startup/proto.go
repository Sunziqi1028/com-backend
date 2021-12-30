package startup

import (
	"ceres/pkg/model"
)

type Mode string

const (
	ModeNONE Mode = "NONE"
	ModeESG  Mode = "ESG"
	ModeNGO  Mode = "NGO"
	ModeDAO  Mode = "DAO"
	ModeCOM  Mode = "COM"
)

type Startup struct {
	model.Base
	ComerID              uint64   `gorm:"comer_id" json:"comerID"`
	Name                 string   `gorm:"name" json:"name"`
	Mode                 Mode     `gorm:"mode" json:"mode"`
	Logo                 string   `gorm:"logo" json:"logo"`
	Mission              string   `gorm:"mission" json:"mission"`
	TokenContractAddress string   `gorm:"token_contract_address" json:"tokenContractAddress"`
	IsSet                bool     `gorm:"is_set" json:"isSet"`
	Wallets              []Wallet `json:"wallets"`
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
