/**
 * @Author: Sun
 * @Description:
 * @File:  request
 * @Version: 1.0.0
 * @Date: 2022/6/29 11:46
 */

package bounty

type BountyRequest struct {
	BountyDetail `json:"bountyDetail"  binding:"required"`
	PayDetail    `json:"payDetail"`
	Deposit      `json:"deposit"  binding:"required"`
	ChainInfo    `json:"chainInfo"  binding:"required"`
}

type BountyDetail struct {
	StartupID         uint64    `json:"startupID"  binding:"required"`
	ComerID           uint64    `json:"comerID"  binding:"required"`
	Title             string    `json:"title"  binding:"required"`
	ExpiresIn         string    `json:"expiresIn"  binding:"required"`
	Contacts          []Contact `json:"contact"  binding:"required"`
	DiscussionLink    string    `json:"discussionLink"`
	ApplicantsSkills  []string  `json:"applicantsSkills"  binding:"required"`
	ApplicantsDeposit int       `json:"applicantsDeposit"`
	Description       string    `json:"description"  binding:"required"`
}

type Contact struct {
	ContactType    uint8  `json:"contactType" binding:"required"` // 1:Email 2:Discord 3:Telegram
	ContactAddress string `json:"contactAddress" binding:"required"`
}

type PayDetail struct {
	Stages []StageType `json:"stages,omitempty"`
	Period PeriodType  `json:"period,omitempty"`
}

type StageType struct {
	SeqNum       int    `json:"id,omitempty"`
	Token1Symbol string `json:"token1Symbol,omitempty"`
	Token1Amount int    `json:"token1Amount,omitempty"`
	Token2Symbol string `json:"token2Symbol,omitempty"`
	Token2Amount int    `json:"token2Amount,omitempty"`
	Terms        string `json:"terms,omitempty"`
}

type PeriodType struct {
	PeriodType   uint8  `json:"periodType,omitempty"` // 1:Days 2:Weeks 3:Months
	PeriodAmount int    `json:"periodAmount,omitempty"`
	HoursPerDay  int    `json:"HoursPerDay,omitempty"`
	Token1Symbol string `json:"token1Symbol,omitempty"`
	Token1Amount int    `json:"token1Amount,omitempty"`
	Token2Symbol string `json:"token2Symbol,omitempty"`
	Token2Amount int    `json:"token2Amount,omitempty"`
	Target       string `json:"target,omitempty"`
}

type Deposit struct {
	TokenSymbol string `json:"tokenSymbol" binding:"required"`
	TokenAmount int    `json:"tokenAmount"`
}

type ChainInfo struct {
	ChainID uint64 `json:"chainID" binding:"required"`
	TxHash  string `json:"txHash" binding:"required"`
}
