/**
* @Author: Sun
* @Description:
* @File:  ether
* @Version: 1.0.0
* @Date: 2022/7/3 10:39
 */

package ether

import (
	"ceres/pkg/initialization/mysql"
	modelTransaction "ceres/pkg/model/transaction"
	serviceTransaction "ceres/pkg/service/transaction"
	"context"
	"github.com/gotomicro/ego/task/ecron"
	"github.com/qiniu/x/log"
	"time"
)

func GetAllContractAddresses() ecron.Ecron {
	job := func(ctx context.Context) error {
		transactions, err := modelTransaction.GetTransaction(mysql.DB)
		log.Info("####GET ALL TRANSACTION_BY_STATUS:", transactions)
		if err != nil {
			return err
		}
		for _, transaction := range transactions {
			contractAddress, status := serviceTransaction.GetContractAddress(transaction.ChainID, transaction.TxHash)
			time.Sleep(5 * time.Second)
			log.Info("the contractAddress is:", contractAddress, "the status is :", status)
			go serviceTransaction.UpdateBountyContractAndTransactoinStatus(mysql.DB, transaction.SourceID, status, contractAddress)
		}
		return nil
	}
	cron := ecron.Load("ceres.cron").Build(ecron.WithJob(job))
	return cron
}
