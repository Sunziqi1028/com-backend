/**
 * @Author: Sun
 * @Description:
 * @File:  request
 * @Version: 1.0.0
 * @Date: 2022/6/29 11:46
 */

package bounty

type BountyRequest struct {
	BountyDetail `json:"bountyDetail"`
	PayDetail    `json:"PayDetail"`
	Deposit      `json:"deposit"`
	ChainInfo    `json:"chainInfo"`
	//status       int `json:"status"`
}

type BountyDetail struct {
	StartupID         uint64    `json:"startupID"`
	ComerID           uint64    `json:"comerID"`
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
	SeqNum       int    `json:"id,omitempty"`
	Token1Symbol string `json:"token1Symbol,omitempty"`
	Token1Amount int    `json:"token1Amount,omitempty"`
	Token2Symbol string `json:"token2Symbol,omitempty"`
	Token2Amount int    `json:"token2Amount,omitempty"`
	Terms        string `json:"terms,omitempty"`
	//Status       int    `json:"status,omitempty"` // unpaid paid
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
	ChainID uint64 `json:"chainID"`
	TxHash  string `json:"txHash"`
}
