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
	CreatedAt         time.Time `json:"createdAt"`
}

type BountyPaymentInfo struct {
	PaymentMode      int `json:"paymentMode" gorm:"payment_mode"`
	FounderDeposit   int `json:"founderDeposit" gorm:"founder_deposit"`
	ApplicantDeposit int `json:"applicantDeposit" gorm:"applicant_deposit"`
}
type PaymentResponse struct {
	BountyPaymentInfo      `json:"bountyPaymentInfo"`
	Rewards                BountyReward `json:"rewards,omitempty"`
	ApplicantsTotalDeposit int          `json:"applicantsTotalDeposit"`
	StageTerms             []StageTerm  `json:"stageTerms"`
	PeriodTerms            `json:"periodTerms"`
	BountyDepositStatus    int `json:"bountyDepositStatus" gorm:"status"`
}

type BountyReward struct {
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1Symbol,omitempty"`
	Token1Amount int    `gorm:"column:token1_amount" json:"token1Amount,omitempty"`
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2Symbol,omitempty"`
	Token2Amount int    `gorm:"column:token2_amount" json:"token2Amount,omitempty"`
}

type StageTerm struct {
	SeqNum       int    `json:"seqNum" gorm:"seq_Num"`
	Status       int    `json:"status" gorm:"status"`
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1Symbol,omitempty"`
	Token1Amount int    `gorm:"column:token1_amount" json:"token1Amount,omitempty"`
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2Symbol,omitempty"`
	Token2Amount int    `gorm:"column:token2_amount" json:"token2Amount,omitempty"`
	Terms        string `json:"terms" gorm:"terms"`
}

type PeriodTerms struct {
	PeriodModes []PeriodMode `json:"periodModes"`
	Terms       string       `json:"terms"`
	HoursPerDay int          `json:"hoursPerDay" gorm:"hours_per_day"`
	PeriodType  int          `json:"periodType" gorm:"period_type"`
}

type PeriodMode struct {
	SeqNum       int    `json:"seqNum" gorm:"seq_Num"`
	Status       int    `json:"status" gorm:"status"`
	Token1Symbol string `gorm:"column:token1_symbol" json:"token1Symbol,omitempty"`
	Token1Amount int    `gorm:"column:token1_amount" json:"token1Amount,omitempty"`
	Token2Symbol string `gorm:"column:token2_symbol" json:"token2Symbol,omitempty"`
	Token2Amount int    `gorm:"column:token2_amount" json:"token2Amount,omitempty"`
}

type PeriodInfo struct {
	HoursPerDay int `json:"hoursPerDay" gorm:"hours_per_day"`
	PeriodType  int `json:"periodType" gorm:"period_type"`
}
type ActivitiesResponse struct {
	ComerID    uint64    `json:"comerID" gorm:"comer_id"`
	Name       string    `json:"name" gorm:"name"`
	Avatar     string    `json:"avatar" gorm:"avatar"`
	SourceType int       `json:"sourceType" gorm:"source_type"`
	Content    string    `json:"content" gorm:"content"`
	Timestamp  time.Time `json:"timestamp" gorm:"timestamp"`
}

//type ActivityContent struct {
//	SourceType int       `json:"sourceType" gorm:"source_type"`
//	Content    string    `json:"content" gorm:"content"`
//	CreatedAt  time.Time `json:"createdAt" gorm:"created_at"`
//}

type ComerInfo struct {
	ComerID    uint64 `json:"comerID"`
	ComerName  string `json:"comerName" gorm:"name"`
	ComerImage string `json:"comerImage" gorm:"avatar"`
}

type StartupListResponse struct {
	Title         string   `gorm:"name" json:"title"`
	Mode          int      `gorm:"mode" json:"mode"`
	Logo          string   `gorm:"logo" json:"logo"`
	ChainID       uint64   `gorm:"chain_id" json:"chainID"`
	TxHash        string   `gorm:"tx_hash" json:"blockChainAddress"`
	ContractAudit string   `gorm:"contract_audit" json:"contractAudit"`
	Website       string   `gorm:"website" json:"website"`
	Discord       string   `gorm:"discord" json:"discord"`
	Twitter       string   `gorm:"twitter" json:"twitter"`
	Telegram      string   `gorm:"telegram" json:"telegram"`
	Docs          string   `json:"docs" gorm:"docs"`
	Mission       string   `json:"mission" gorm:"mission"`
	Tag           []string `json:"tag"`
}

type BountyApplicantsResponse struct {
	Applicants []Applicant
}

type Applicant struct {
	ComerID     uint64    `json:"comerID"`
	Image       string    `json:"image"`
	Name        string    `json:"name"`
	Description string    `json:"desription"`
	ApplyAt     time.Time `json:"applyAt"`
}

type FounderResponse struct {
	ComerID          uint64   `json:"comerID"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	ApplicantsSkills []string `json:"applicantsSkills"`
	TimeZone         string   `json:"timeZone"`
	Location         string   `gorm:"column:location" json:"location"`
	Email            string   `gorm:"column:email" json:"email"`
}

type ApprovedResponse struct {
	ComerID          uint64   `json:"comerID"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	ApplicantsSkills []string `json:"applicantsSkills"`
}

type DepositRecordsResponse struct {
	DepositRecords []DepositRecord
}

type DepositRecord struct {
	ComerID     uint64    `json:"comerID"`
	Name        string    `json:"name"`
	Time        time.Time `json:"time"`
	TokenAmount int       `json:"tokenAmount" gorm:"token_Amount"`
	Access      int       `json:"access"`
}
