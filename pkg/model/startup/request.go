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
	KYC           *string  `json:"kyc"`
	ContractAudit *string  `json:"contractAudit"`
	HashTags      []string `json:"hashTags"`
	Website       *string  `json:"website"`
	Discord       *string  `json:"discord"`
	Twitter       *string  `json:"twitter"`
	Telegram      *string  `json:"telegram"`
	Docs          *string  `json:"docs"`
}

type UpdateStartupFinanceSettingRequest struct {
	TokenContractAddress *string `json:"tokenContractAddress" binding:"required"`
	LaunchNetwork        *int    `json:"launchNetwork" binding:"required"`
	TokenName            *string `json:"tokenName" binding:"required"`
	TokenSymbol          *string `json:"tokenSymbol" binding:"required"`
	TotalSupply          *int64  `json:"totalSupply" binding:"required"`
	PresaleStart         *string `json:"presaleStart" binding:"required"`
	PresaleEnd           *string `json:"presaleEnd" binding:"required"`
	LaunchDate           *string `json:"launchDate" binding:"required"`
	Wallets              []struct {
		WalletName    string `json:"walletName" binding:"min=1,max=50"`
		WalletAddress string `json:"walletAddress" binding:"len=42,startswith=0x"`
	} `json:"wallets" binding:"required"`
}
