/**
 * @Author: Sun
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2022/6/29 13:17
 */

package bounty

type GetStartupsResponse struct {
	StartupID uint64 `gorm:"column:id" json:"startupID"`
	Name      string `gorm:"column:name" json:"name"`
}

type CreateBountyResponse struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

type ContractInfoResponse struct {
	ContractAddress string
	Status          uint64
}

type GetTransactions struct {
	ChainID  uint64 `gorm:"column:chain_id"`
	TxHash   string `gorm:"column:tx_hash"`
	SourceID uint64 `gorm:"column:source_id"`
}
