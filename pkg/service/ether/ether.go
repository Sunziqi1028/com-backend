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

//transactions, err := modelTransaction.GetTransaction(mysql.DB)
//fmt.Println(&transactions) // 注释
//if err != nil {
//	return
//}
//for _, transaction := range transactions {
//	//var contractChan = make(chan *modelBonuty.ContractInfoResponse, 1024)
//	contractAddress, status := serviceTransaction.GetContractAddress(transaction.ChainID, transaction.TxHash)
//	time.Sleep(5 * time.Second)
//	go serviceTransaction.UpdateBountyContractAndTransactoinStatus(mysql.DB, transaction.SourceID, status, contractAddress)

//contractInfo := &modelBonuty.ContractInfoResponse{
//	ContractAddress: contractAddress,
//	Status:          status,
//}
//select {
//case contractChan <- contractInfo:
//	for contract := range contractChan {
//		serviceTransaction.UpdateBountyContractAndTransactoinStatus(mysql.DB, transaction.SourceID, contract.Status, contract.ContractAddress)
//		return
//	}
//case <-time.After(5 * time.Second):
//	fmt.Println("get contract address time over!")
//	return
//}
//}
//
//time.Sleep(30 * time.Second)
//	}
//}()
//return
//}
