package startup

import (
	"ceres/pkg/model"
	"ceres/pkg/model/tag"
	"database/sql"
	"encoding/json"
	"time"
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
	TxHash               string    `gorm:"tx_hash" json:"blockChainAddress"`
	IsSet                bool      `gorm:"is_set" json:"isSet"`
	KYC                  string    `gorm:"kyc" json:"kyc"`
	ContractAudit        string    `gorm:"contract_audit" json:"contractAudit"`
	HashTags             []tag.Tag `gorm:"many2many:tag_target_rel;foreignKey:ID;joinForeignKey:TargetID;" json:"hashTags"`
	Website              string    `gorm:"website" json:"website"`
	Discord              string    `gorm:"discord" json:"discord"`
	Twitter              string    `gorm:"twitter" json:"twitter"`
	Telegram             string    `gorm:"telegram" json:"telegram"`
	Docs                 string    `gorm:"docs" json:"docs"`
	PresaleDate          NullTime  `gorm:"presale_date" json:"presaleDate"`
	LaunchDate           NullTime  `gorm:"launch_date" json:"launchDate"`
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

// BasicSetting Startup security and social setting
type BasicSetting struct {
	KYC           string `gorm:"kyc" json:"kyc"`
	ContractAudit string `gorm:"contract_audit" json:"contractAudit"`
	Website       string `gorm:"website" json:"website"`
	Discord       string `gorm:"discord" json:"discord"`
	Twitter       string `gorm:"twitter" json:"twitter"`
	Telegram      string `gorm:"telegram" json:"telegram"`
	Docs          string `gorm:"docs" json:"docs"`
}

// FinanceSetting Startup finance setting
type FinanceSetting struct {
	TokenContractAddress string    `gorm:"token_contract_address" json:"tokenContractAddress"`
	PresaleDate          time.Time `gorm:"presale_date" json:"presaleDate"`
	LaunchDate           time.Time `gorm:"launch_date" json:"launchDate"`
	//Wallets              []Wallet  `json:"wallets"`
}

type NullTime struct {
	sql.NullTime
}

func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal("")
	}
}

func (v *NullTime) UnmarshalJSON(data []byte) error {
	var s *time.Time
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.Time = *s
	} else {
		v.Valid = false
	}
	return nil
}
