package startup

import (
	"ceres/pkg/model"
)

type ListStartupRequest struct {
	model.ListRequest
	Keyword string `form:"keyword"`
	Mode    Mode   `form:"mode"`
}

type UpdateStartupBasicSettingRequest struct {
	KYC           *string  `json:"kyc" binding:"required"`
	ContractAudit *string  `json:"contractAudit" binding:"required"`
	HashTags      []string `json:"hashTags" binding:"required"`
	Website       *string  `json:"website" binding:"required"`
	Discord       *string  `json:"discord" binding:"required"`
	Twitter       *string  `json:"twitter" binding:"required"`
	Telegram      *string  `json:"telegram" binding:"required"`
	Docs          *string  `json:"docs" binding:"required"`
}

type UpdateStartupFinanceSettingRequest struct {
	TokenContractAddress *string `json:"tokenContractAddress" binding:"required"`
	LaunchNetwork        *int    `json:"launchNetwork" binding:"required"`
	PresaleStart         *string `json:"presaleStart" binding:"required"`
	PresaleEnd           *string `json:"presaleEnd" binding:"required"`
	LaunchDate           *string `json:"launchDate" binding:"required"`
	Wallets              []struct {
		WalletName    string `json:"walletName" binding:"min=1,max=50"`
		WalletAddress string `json:"walletAddress" binding:"len=42,startswith=0x"`
	} `json:"wallets" binding:"required"`
}
