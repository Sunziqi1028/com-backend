/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/6/28 09:02
 */

package bounty

import "time"

type CreatBountyInfo struct {
	BountyDetail `json:"bountyDetail"`
	PayDetail    `json:"PayDetail"`
	Deposit      `json:"deposit"`
	ChainInfo    `json:"chainInfo"`
	status       int `json:"status"`
}

type BountyDetail struct {
	StartupID         string    `json:"startupID"`
	ComerID           int64     `json:"comerID"`
	Title             string    `json:"title"`
	ExpiresIn         string    `json:"expiresIn"`
	Contacts          []Contact `json:"contact"`
	DiscussionLink    string    `json:"discussionLink"`
	ApplicantsSkills  []string  `json:"applicantsSkills"`
	ApplicantsDeposit int       `json:"applicantsDeposit"`
	Description       string    `json:"description"`
}

type Contact struct {
	Email    string `json:"email,omitempty"`
	Telegram string `json:"telegram,omitempty"`
	Discord  string `json:"discord,omitempty"`
}

type PayDetail struct {
	Stages []StageType `json:"stage,omitempty"`
	Period PeriodType  `json:"period,omitempty"`
}

type StageType struct {
	Id           string `json:"id,omitempty"`
	Token1Symbol string `json:"token1Symbol,omitempty"`
	Token1Amount int    `json:"token1Amount,omitempty"`
	Token2Symbol string `json:"token2Symbol,omitempty"`
	Token2Amount int    `json:"token2Amount,omitempty"`
	Terms        string `json:"terms,omitempty"`
	Status       string `json:"status,omitempty"` // unpaid paid
}

type PeriodType struct {
	PeriodType   string `json:"days,weeks,months,omitempty"`
	HoursPerDay  int    `json:"HoursPerDay,omitempty"`
	Token1Symbol string `json:"token1Symbol,omitempty"`
	Token1Amount int    `json:"token1Amount,omitempty"`
	Token2Symbol string `json:"token2Symbol,omitempty"`
	Token2Amount int    `json:"token2Amount,omitempty"`
	Target       string `json:"target"`
}

type Deposit struct {
	TokenSymbol string `json:"tokenSymbol,omitempty"`
	TokenAmount int    `json:"tokenAmount,omitempty"`
}

type ChainInfo struct {
	ChainID string `json:"chainID"`
	TxHash  string `json:"txHash"`
}

type BountyModel struct {
	ChainID            int64     `json:"chainID"`
	TxHash             string    `json:"txHash"`
	DepositContract    string    `json:"depositContract"`
	StartupID          int64     `json:"startupID"`
	ComerID            int64     `json:"comerID"`
	Title              string    `json:"title"`
	ApplyCutoffDate    time.Time `json:"applyCutoffDate"`
	DiscussionLink     string    `json:"discussionLink"`
	DepositTokenSymbol string    `json:"depositTokenSymbol"`
	ApplicationDeposit int       `json:"applicationDeposit"`
	FounderDeposit     int       `json:"founderDeposit"`
	Description        string    `json:"description"`
	PaymentMode        int       `json:"paymentMode"`
	Status             int       `json:"status"`
	TotalRewardToken   int       `json:"totalRewardToken"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	IsDeleted          int `json:"isDeleted"`
}

type ApplicantModel struct {
	BountyID    int64     `json:"bountyID"`
	ComerID     int64     `json:"comerID"`
	ApplyAt     time.Time `json:"applyAt"`
	RevokeAt    time.Time `json:"revokeAt"`
	ApproveAt   time.Time `json:"approveAt"`
	QuitAt      time.Time `json:"quitAt"`
	SubmitAt    time.Time `json:"submitAt"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ContactModel struct {
	BountyID       int64  `json:"bountyID"`
	ContactType    int    `json:"contactType"`
	ContactAddress string `json:"contactAddress"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type DepositModel struct {
	ChainID     int64  `json:"chainID"`
	TxHash      string `json:"txHash"`
	Status      int    `json:"status"`
	BountyID    int64  `json:"bountyID"`
	ComerID     int64  `json:"comerID"`
	Access      int    `json:"access"`
	TokenSymbol string `json:"tokenSymbol,omitempty"`
	TokenAmount int    `json:"tokenAmount,omitempty"`
	TimeStamp   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaymentPeriodModel struct {
	BountyID     int64  `json:"bountyID"`
	PeriodType   int    `json:"periodType"`
	PeriodAmount int64  `json:"periodAmount"`
	HoursPerDay  int    `json:"hoursPerDay"`
	Token1Symbol string `json:"token1Symbol,omitempty"`
	Token1Amount int    `json:"token1Amount,omitempty"`
	Token2Symbol string `json:"token2Symbol,omitempty"`
	Token2Amount int    `json:"token2Amount,omitempty"`
	Target       string `json:"target"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PaymentTermsModel struct {
	BountyID     int64  `json:"bountyID"`
	PaymentMode  int    `json:"paymentMode"`
	Token1Symbol string `json:"token1Symbol,omitempty"`
	Token1Amount int    `json:"token1Amount,omitempty"`
	Token2Symbol string `json:"token2Symbol,omitempty"`
	Token2Amount int    `json:"token2Amount,omitempty"`
	Terms        string `json:"terms"`
	SeqNum       int    `json:"seqNum"`
	Status       int    `json:"status"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type TransactionModel struct {
	ChainID    int64  `json:"chainID"`
	TxHash     string `json:"txHash"`
	timeStamp  time.Time
	Status     int   `json:"status,omitempty"`
	SourceType int   `json:"sourceType"`
	SourceID   int64 `json:"sourceID"`
	RetryTimes int   `json:"retryTimes"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type PostUpdateModel struct {
	SourceType int       `gorm:"sourceType"`
	SourceID   int64     `gorm:"sourceID"`
	ComerID    int64     `gorm:"comerID"`
	TimeStamp  time.Time // post time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
