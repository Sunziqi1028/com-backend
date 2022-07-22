package crowdfunding

import (
	"ceres/pkg/model"
	"encoding/json"
	"github.com/shopspring/decimal"
	"time"
)

type CrowdfundingStatus int

const (
	Pending CrowdfundingStatus = iota
	Upcoming
	Live
	Ended
	Cancelled
	OnChainFailure
)

type ChainInfo struct {
	ChainId uint64 `gorm:"column:chain_id;unique_index:chain_tx_uindex" json:"chainId"`
	TxHash  string `gorm:"tx_hash;unique_index:chain_tx_uindex" json:"txHash"`
}

type Crowdfunding struct {
	model.Base
	ChainInfo
	SellInfo
	BuyInfo
	// update by querying chain???
	CrowdfundingContract string          `gorm:"crowdfunding_contract" json:"crowdfundingContract,omitempty"`
	StartupId            uint64          `gorm:"startup_id" json:"startupId"`
	ComerId              uint64          `gorm:"comer_id" json:"comerId"`
	RaiseGoal            decimal.Decimal `gorm:"raise_goal" json:"raiseGoal"`
	// update by querying chain???
	RaiseBalance decimal.Decimal `gorm:"raise_balance" json:"raiseBalance"`

	TeamWallet  string             `gorm:"team_wallet" json:"teamWallet"`
	SwapPercent decimal.Decimal    `gorm:"swap_percent" json:"swapPercent"`
	StartTime   time.Time          `gorm:"start_time" json:"startTime"`
	EndTime     time.Time          `gorm:"end_time" json:"endTime"`
	Poster      string             `gorm:"poster" json:"poster"`
	Youtube     string             `gorm:"youtube" json:"youtube"`
	Detail      string             `gorm:"detail" json:"detail"`
	Description string             `gorm:"description" json:"description"`
	Status      CrowdfundingStatus `gorm:"status" json:"status"`
}

func (c Crowdfunding) Json() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type SellInfo struct {
	SellTokenContract string          `gorm:"sell_token_contract" json:"sellTokenContract"`
	SellTokenName     string          `gorm:"sell_token_name" json:"sellTokenName,omitempty"`
	SellTokenSymbol   string          `gorm:"sell_token_symbol" json:"sellTokenSymbol,omitempty"`
	SellTokenDecimals int             `gorm:"sell_token_decimals" json:"sellTokenDecimals,omitempty"`
	SellTokenSupply   decimal.Decimal `gorm:"sell_token_supply" json:"sellTokenSupply,omitempty"`
	SellTokenDeposit  decimal.Decimal `gorm:"sell_token_deposit" json:"sellTokenDeposit"`
	SellTokenBalance  decimal.Decimal `gorm:"sell_token_balance" json:"sellTokenBalance"`
	MaxSellPercent    decimal.Decimal `gorm:"max_sell_percent" json:"maxSellPercent"`
	SellTax           decimal.Decimal `gorm:"sell_tax" json:"sellTax"`
}

type BuyInfo struct {
	BuyTokenContract string          `gorm:"buy_token_contract" json:"buyTokenContract"`
	BuyTokenName     string          `gorm:"buy_token_name" json:"buyTokenName,omitempty"`
	BuyTokenSymbol   string          `gorm:"buy_token_symbol" json:"buyTokenSymbol,omitempty"`
	BuyTokenDecimals int             `gorm:"buy_token_decimals" json:"buyTokenDecimals,omitempty"`
	BuyTokenSupply   decimal.Decimal `gorm:"buy_token_supply" json:"buyTokenSupply,omitempty"`
	BuyPrice         decimal.Decimal `gorm:"buy_price" json:"buyPrice"`
	MaxBuyAmount     decimal.Decimal `gorm:"max_buy_amount" json:"maxBuyAmount"`
}

func (c Crowdfunding) TableName() string {
	return "crowdfunding"
}

type CrowdfundingSwapStatus int
type SwapAccess int

const (
	PendingSwap CrowdfundingSwapStatus = iota
	Success
	Failure
)
const (
	Invest SwapAccess = iota + 1
	Withdraw
)

type CrowdfundingSwap struct {
	model.RelationBase
	ChainInfo
	Timestamp       time.Time              `gorm:"timestamp" json:"timestamp"`
	Status          CrowdfundingSwapStatus `gorm:"status" json:"status"`
	CrowdfundingId  uint64                 `gorm:"crowdfunding_id" json:"crowdfundingId"`
	ComerId         uint64                 `gorm:"comer_id" json:"comerId"`
	Access          SwapAccess             `gorm:"access" json:"access"`
	BuyTokenSymbol  string                 `gorm:"buy_token_symbol" json:"buyTokenSymbol"`
	BuyTokenAmount  decimal.Decimal        `gorm:"buy_token_amount" json:"buyTokenAmount"`
	SellTokenSymbol string                 `gorm:"sell_token_symbol" json:"sellTokenSymbol"`
	SellTokenAmount decimal.Decimal        `gorm:"sell_token_amount" json:"sellTokenAmount"`
	Price           decimal.Decimal        `gorm:"price" json:"price"`
}

func (c CrowdfundingSwap) TableName() string {
	return "crowdfunding_swap"
}
