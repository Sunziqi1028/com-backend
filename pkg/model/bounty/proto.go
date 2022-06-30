/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/6/28 09:02
 */

package bounty

import (
	"time"
)

type BountyModel struct {
	ID                 uint64    `gorm:"column:id" json:"id"`
	ChainID            uint64    `gorm:"column:chain_id" json:"chainID"`
	TxHash             string    `gorm:"column:tx_hash" json:"txHash"`
	DepositContract    string    `gorm:"column:deposit_contract" json:"depositContract"`
	StartupID          uint64    `gorm:"column:startup_id" json:"startupID"`
	ComerID            uint64    `gorm:"column:comer_id" json:"comerID"`
	Title              string    `gorm:"column:title" json:"title"`
	ApplyCutoffDate    time.Time `gorm:"column:apply_cutoff_date" json:"applyCutoffDate"`
	DiscussionLink     string    `gorm:"column:discussion_link" json:"discussionLink"`
	DepositTokenSymbol string    `gorm:"column:deposit_token_symbol" json:"depositTokenSymbol"`
	ApplicantDeposit   int       `gorm:"column:applicant_deposit" json:"applicationDeposit"`
	FounderDeposit     int       `gorm:"column:founder_deposit" json:"founderDeposit"`
	Description        string    `gorm:"column:description" json:"description"`
	PaymentMode        int       `gorm:"column:payment_mode" json:"paymentMode"`
	Status             int       `gorm:"column:status" json:"status"`
	TotalRewardToken   int       `gorm:"column:total_reward_token" json:"totalRewardToken"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updatedAt"`
	IsDeleted          int       `gorm:"column:is_deleted" gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName the BountyModel table name for gorm
func (BountyModel) TableName() string {
	return "bounty"
}

type ApplicantModel struct {
	BountyID    uint64    `gorm:"column:bounty_id" json:"bountyID"`
	ComerID     uint64    `gorm:"column:comer_id" json:"comerID"`
	ApplyAt     time.Time `gorm:"column:apply_at" json:"applyAt"`
	RevokeAt    time.Time `gorm:"column:revoke_at" json:"revokeAt"`
	ApproveAt   time.Time `gorm:"column:approve_at" json:"approveAt"`
	QuitAt      time.Time `gorm:"column:quit_at" json:"quitAt"`
	SubmitAt    time.Time `gorm:"column:submit_at" json:"submitAt"`
	Status      int       `gorm:"column:status" json:"status"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

// TableName the ApplicantModel table name for gorm
func (ApplicantModel) TableName() string {
	return "bounty_applicant"
}

type ContactModel struct {
	BountyID       uint64    `gorm:"column:bounty_id" json:"bountyID"`
	ContactType    int       `gorm:"column:contact_type" json:"contactType"`
	ContactAddress string    `gorm:"column:contact_address" json:"contactAddress"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

// TableName the ContactModel table name for gorm
func (ContactModel) TableName() string {
	return "bounty_contact"
}

type DepositModel struct {
	ChainID     uint64    `gorm:"column:chain_id" json:"chainID"`
	TxHash      string    `gorm:"column:tx_hash" json:"txHash"`
	Status      int       `gorm:"column:status" json:"status"`
	BountyID    uint64    `gorm:"column:bounty_id" json:"bountyID"`
	ComerID     uint64    `gorm:"column:comer_id" json:"comerID"`
	Access      int       `gorm:"column:access" json:"access"`
	TokenSymbol string    `gorm:"column:token_symbol" json:"tokenSymbol,omitempty"`
	TokenAmount int       `gorm:"column:token_amount" json:"tokenAmount,omitempty"`
	TimeStamp   time.Time `gorm:"column:timestamp"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

// TableName the DepositModel table name for gorm
func (DepositModel) TableName() string {
	return "bounty_deposit"
}

type PaymentPeriodModel struct {
	BountyID     uint64    `gorm:"column:bounty_id" json:"bountyID"`
	PeriodType   int       `gorm:"column:period_type" json:"periodType"`
	PeriodAmount int64     `gorm:"column:period_aomunt" json:"periodAmount"`
	HoursPerDay  int       `gorm:"column:hours_per_day" json:"hoursPerDay"`
	Token1Symbol string    `gorm:"column:token1_symbol" json:"token1Symbol,omitempty"`
	Token1Amount int       `gorm:"column:token1_amount" json:"token1Amount,omitempty"`
	Token2Symbol string    `gorm:"column:token2_symbol" json:"token2Symbol,omitempty"`
	Token2Amount int       `gorm:"column:token2_amount" json:"token2Amount,omitempty"`
	Target       string    `gorm:"column:target" json:"target"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

// TableName the PaymentPeriodModel table name for gorm
func (PaymentPeriodModel) TableName() string {
	return "bounty_payment_period"
}

type PaymentTermsModel struct {
	BountyID     uint64    `gorm:"column:bounty_id" json:"bountyID"`
	PaymentMode  int       `gorm:"column:payment_mode" json:"paymentMode"`
	Token1Symbol string    `gorm:"column:token1_symbol" json:"token1Symbol,omitempty"`
	Token1Amount int       `gorm:"column:token1_amount" json:"token1Amount,omitempty"`
	Token2Symbol string    `gorm:"column:token2_symbol" json:"token2Symbol,omitempty"`
	Token2Amount int       `gorm:"column:token2_amount" json:"token2Amount,omitempty"`
	Terms        string    `gorm:"column:terms" json:"terms"`
	SeqNum       int       `gorm:"column:seq_num" json:"seqNum"`
	Status       int       `gorm:"column:status" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

// TableName the PaymentTermsModel table name for gorm
func (PaymentTermsModel) TableName() string {
	return "bounty_payment_terms"
}

type TransactionModel struct {
	ChainID    uint64    `gorm:"column:chain_id" json:"chainID"`
	TxHash     string    `gorm:"column:tx_hash" json:"txHash"`
	TimeStamp  time.Time `gorm:"column:timestamp"`
	Status     int       `gorm:"column:status" json:"status,omitempty"` // 0:Pending 1:Success 2:Failure
	SourceType int       `gorm:"column:source_type" json:"sourceType"`
	SourceID   int64     `gorm:"column:source_id" json:"sourceID"`
	RetryTimes int       `gorm:"column:retry_times" json:"retryTimes"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

// TableName the TransactionModel table name for gorm
func (TransactionModel) TableName() string {
	return "transaction"
}

type PostUpdateModel struct {
	SourceType int       `gorm:"sourceType"`
	SourceID   uint64    `gorm:"sourceID"`
	ComerID    uint64    `gorm:"comerID"`
	Content    string    `gorm:"column:content"`
	TimeStamp  time.Time `gorm:"column:timestamp"` // post time
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

// TableName the PostUpdateModel table name for gorm
func (PostUpdateModel) TableName() string {
	return "post_update"
}
