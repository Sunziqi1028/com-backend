/**
 * @Author: Sun
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2022/7/3 13:58
 */

package transaction

type GetTransactions struct {
	ChainID  uint64 `gorm:"column:chain_id"`
	TxHash   string `gorm:"column:tx_hash"`
	SourceID uint64 `gorm:"column:source_id"`
}
