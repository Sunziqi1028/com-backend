/**
 * @Author: Sun
 * @Description:
 * @File:  transaction
 * @Version: 1.0.0
 * @Date: 2022/7/3 10:37
 */

package transaction

import (
	"ceres/pkg/initialization/eth"
	"ceres/pkg/model/bounty"
	model "ceres/pkg/model/transaction"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/qiniu/x/log"
	"gorm.io/gorm"
	"time"
)

const (
	Pending                      = 0
	Success                      = 1
	Failure                      = 2
	BountyDepositContractCreated = 1
	BountyDepositAccount         = 2
	ReceiptSuccess               = 1
	ReceiptFailure               = 0
)

func CreateTransaction(db *gorm.DB, bountyID uint64, request *bounty.BountyRequest) error {
	transaction := &model.Transaction{
		ChainID:    request.ChainID,
		TxHash:     request.TxHash,
		TimeStamp:  time.Now(),
		Status:     Pending,
		SourceType: BountyDepositContractCreated,
		RetryTimes: 1,
		SourceID:   int64(bountyID),
	}
	if err := model.CreateTransaction(db, transaction); err != nil {
		return err
	}
	return nil
}

func UpdateBountyContractAndTransactoinStatus(tx *gorm.DB, bountyID, status uint64, contractAddress string) {
	err := model.UpdateTransactionStatus(tx, bountyID, status)
	if err != nil {
		log.Warn(err)
	}
	err = bounty.UpdateBountyDepositContract(tx, bountyID, contractAddress)
	if err != nil {
		log.Warn(err)
	}
	err = bounty.UpdateBountyDepositStatus(tx, bountyID, status)
	if err != nil {
		log.Warn(err)
	}
}

func GetContractAddress(chainID uint64, txHashString string) (contractAddress string, status uint64) {
	txHash := common.HexToHash(txHashString)
	tx, isPending, err := eth.Client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Warn(err)
		return "", Failure
	}
	if isPending == false {
		receipt, err := eth.Client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Warn(err)
			return "", Failure
		}
		if receipt.Status == ReceiptFailure {
			return "", Failure
		}

		contractAddress = receipt.ContractAddress.String()

		return contractAddress, Success
	}
	return "", Pending
}
