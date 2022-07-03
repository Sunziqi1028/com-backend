/**
 * @Author: Sun
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2022/6/29 13:17
 */

package bounty


type ContractInfoResponse struct {
	ContractAddress string
	Status          uint64
}


type DetailItem struct {
	Logo                string    `json:"logo"`
	Title               string    `json:"title"`
	Status              string    `json:"status"`
	PaymentType         string    `json:"paymentType"`
	CreatedTime         time.Time `json:"createdTime"`
	Rewards             []string  `json:"rewards"`
	ApproveCount        int       `json:"approveCount"`
	ApplicationSkills   []string  `json:"applicationSkills"`
	DepositRequirements []string  `json:"depositRequirements"`
}

type TabListResponse struct {
	BountyCount int `json:"bountyCount"`
	PageParam
	TotalPages int `json:"totalPages"`
	Records    []*DetailItem
}
