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
	PayDetail    `json:"payDetail"  binding:"required"`
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
	ApplicantsDeposit int       `json:"applicantsDeposit"  binding:"required"`
	Description       string    `json:"description"  binding:"required"`
}

type Contact struct {
	ContactType    uint8  `json:"contactType" binding:"required"` // 1:Email 2:Discord 3:Telegram
	ContactAddress string `json:"contactAddress" binding:"required"`
}

type PayDetail struct {
	Stages []StageType `json:"stage,omitempty"`
	Period PeriodType  `json:"period,omitempty"`
}

type StageType struct {
	SeqNum       int    `json:"id" binding:"required"`
	Token1Symbol string `json:"token1Symbol" binding:"required"`
	Token1Amount int    `json:"token1Amount" binding:"required"`
	Token2Symbol string `json:"token2Symbol" binding:"required"`
	Token2Amount int    `json:"token2Amount" binding:"required"`
	Terms        string `json:"terms" binding:"required"`
}

type PeriodType struct {
	PeriodType   uint8  `json:"periodType" binding:"required"` // 1:Days 2:Weeks 3:Months
	PeriodAmount int    `json:"periodAmount" binding:"required"`
	HoursPerDay  int    `json:"HoursPerDay" binding:"required"`
	Token1Symbol string `json:"token1Symbol" binding:"required"`
	Token1Amount int    `json:"token1Amount" binding:"required"`
	Token2Symbol string `json:"token2Symbol" binding:"required"`
	Token2Amount int    `json:"token2Amount" binding:"required"`
	Target       string `json:"target" binding:"required"`
}

type Deposit struct {
	TokenSymbol string `json:"tokenSymbol" binding:"required"`
	TokenAmount int    `json:"tokenAmount" binding:"required"`
}

type ChainInfo struct {
	ChainID uint64 `json:"chainID" binding:"required"`
	TxHash  string `json:"txHash" binding:"required"`
}
