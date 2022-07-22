/**
 * @Author: Sun
 * @Description:
 * @File:  transaction
 * @Version: 1.0.0
 * @Date: 2022/7/3 10:57
 */

package transaction

import "gorm.io/gorm"

func CreateTransaction(db *gorm.DB, transaction *Transaction) error {
	return db.Create(&transaction).Error
}

// UpdateTransactionStatus update transaction status by corresponding source id
func UpdateTransactionStatus(db *gorm.DB, sourceID uint64, status int) error {
	return db.Model(&Transaction{}).Where("source_id = ?", sourceID).Update("status", status).Error
}

func GetTransaction(db *gorm.DB) (transactionResponse []*GetTransactions, err error) {
	err = db.Table("transaction").Select("chain_id, tx_hash, source_id").Where("status = 0").Find(&transactionResponse).Error
	if err != nil {
		return nil, err
	}
	return transactionResponse, nil
}
