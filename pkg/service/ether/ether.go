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
	modelBonuty "ceres/pkg/model/bounty"
	modelTransaction "ceres/pkg/model/transaction"
	serviceTransaction "ceres/pkg/service/transaction"
	"github.com/qiniu/x/log"
	"time"
)

func GetAllContractAddresses() {
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for {
			t := ticker.C
			log.Infof("time now is :", &t)
			transactions, err := modelTransaction.GetTransaction(mysql.DB)
			log.Infof("transaction: %v", transactions) // 注释
			if err != nil {
				return
			}
			for _, transaction := range transactions {
				var contractChan = make(chan *modelBonuty.ContractInfoResponse, 1)
				contractAddress, status := serviceTransaction.GetContractAddress(transaction.ChainID, transaction.TxHash)
				contractInfo := &modelBonuty.ContractInfoResponse{
					ContractAddress: contractAddress,
					Status:          status,
				}
				select {
				case contractChan <- contractInfo:
					for contract := range contractChan {
						serviceTransaction.UpdateBountyContractAndTransactoinStatus(mysql.DB, transaction.SourceID, contract.Status, contract.ContractAddress)
						return
					}
				case <-time.After(5 * time.Second):
					log.Info("get contract address time over!")
					return
				}
				return
			}
		}
	}()
}
