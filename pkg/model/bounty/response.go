/**
 * @Author: Sun
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2022/6/29 13:17
 */

package bounty

import (
	"time"
)

type ContractInfoResponse struct {
	ContractAddress string
	Status          uint64
}

type DetailItem struct {
	BountyId            uint64    `json:"bountyId"`
	StartupId           uint64    `json:"startupId"`
	Logo                string    `json:"logo"`
	Title               string    `json:"title"`
	Status              string    `json:"status"`
	OnChainStatus       string    `json:"onChainStatus"`
	PaymentType         string    `json:"paymentType"`
	CreatedTime         time.Time `json:"createdTime"`
	Rewards             *[]Reward `json:"rewards"`
	ApplicantCount      int       `json:"applicantCount"`
	ApplicationSkills   []string  `json:"applicationSkills"`
	DepositRequirements int       `json:"depositRequirements"`
}

type TabListResponse struct {
	BountyCount int `json:"bountyCount"`
	PageParam
	TotalPages int `json:"totalPages"`
	Records    []*DetailItem
}

type Reward struct {
	TokenSymbol string `json:"tokenSymbol"`
	Amount      int    `json:"amount"`
}

type DetailResponse struct {
	Title             string    `json:"title"`
	Status            int       `json:"status"`
	ApplicantSkills   []string  `json:"applicantSkills"`
	DiscussionLink    string    `json:"discussionLink"`
	ExpiresIn         string    `json:"expiresIn"`
	ApplicantsDeposit int       `json:"applicantsDeposit"`
	Description       string    `json:"description"`
	Contacts          []Contact `json:"contact"`
	CreatAt           time.Time `json:"creatAt"`
}

type PaymentResponse struct {
	ComerID                uint64       `json:"comerID" gorm:"comer_id"`
	PaymentMode            int          `json:"paymentMode" gorm:"payment_mode"`
	Rewards                BountyReward `json:"rewards"`
	FounderDeposit         int          `json:"founderDeposit" gorm:"founder_deposit"`
	ApplicantsTotalDeposit int          `json:"applicantsTotalDeposit"`
	Terms                  []Term       `json:"terms"`
}

type BountyReward struct {
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1Symbol,omitempty"`
	Token1Amount int    `gorm:"column:token1_amount" json:"token1Amount,omitempty"`
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2Symbol,omitempty"`
	Token2Amount int    `gorm:"column:token2_amount" json:"token2Amount,omitempty"`
}

type Term struct {
	SeqNum       int    `json:"seqNum" gorm:"seq_Num"`
	Status       int    `json:"status" gorm:"status"`
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1Symbol,omitempty"`
	Token1Amount int    `gorm:"column:token1_amount" json:"token1Amount,omitempty"`
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2Symbol,omitempty"`
	Token2Amount int    `gorm:"column:token2_amount" json:"token2Amount,omitempty"`
	Terms        string `json:"terms" gorm:"terms"`
}

type ActivitiesResponse struct {
	ComerInfo
	Activities []Activity
}

type Activity struct {
	Content  string    `json:"content" gorm:"content"`
	CreateAt time.Time `json:"creatAt" gorm:"createAt"`
}

type ComerInfo struct {
	ComerName  string `json:"comerName" gorm:"name"`
	ComerImage string `json:"comerImage" gorm:"avatar"`
}
type StartupListResponse struct {
	Name          string `gorm:"name" json:"name"`
	Mode          int    `gorm:"mode" json:"mode"`
	Logo          string `gorm:"logo" json:"logo"`
	ChainID       uint64 `gorm:"chain_id" json:"chainID"`
	TxHash        string `gorm:"tx_hash" json:"blockChainAddress"`
	ContractAudit string `gorm:"contract_audit" json:"contractAudit"`
	Website       string `gorm:"website" json:"website"`
	Discord       string `gorm:"discord" json:"discord"`
	Twitter       string `gorm:"twitter" json:"twitter"`
	Telegram      string `gorm:"telegram" json:"telegram"`
}

type BountyApplicantsResponse struct {
	Applicants []Applicant
}

type Applicant struct {
	Image       string    `json:"image"`
	Name        string    `json:"name"`
	Description string    `json:"desription"`
	ApplyAt     time.Time `json:"applyAt"`
}

type FounderResponse struct {
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	ApplicantsSkills []string `json:"applicantsSkills"`
	TimeZone         string   `json:"timeZone"`
}

type ApprovedResponse struct {
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	ApplicantsSkills []string `json:"applicantsSkills"`
}

type DepositRecordsResponse struct {
	DepositRecords []DepositRecord
}

type DepositRecord struct {
	Name          string    `json:"name"`
	Time          time.Time `json:"time"`
	DepositAmount int       `json:"depositAmount"`
	Access        int       `json:"access"`
}
