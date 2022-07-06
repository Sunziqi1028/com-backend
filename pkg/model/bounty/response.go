/**
 * @Author: Sun
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2022/6/29 13:17
 */

package bounty

import "time"

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
