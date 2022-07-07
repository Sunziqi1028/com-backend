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
	"fmt"
	"time"
)

func GetAllContractAddresses() {
	//ticker := time.NewTicker(30 * time.Second)
	//
	//for {
	//t := ticker.C
	//fmt.Println("time now is :", &t)
	transactions, err := modelTransaction.GetTransaction(mysql.DB)
	fmt.Println(&transactions) // 注释
	if err != nil {
		return
	}
	for _, transaction := range transactions {
		//var contractChan = make(chan *modelBonuty.ContractInfoResponse, 1024)
		contractAddress, status := serviceTransaction.GetContractAddress(transaction.ChainID, transaction.TxHash)
		time.Sleep(5 * time.Second)
		go serviceTransaction.UpdateBountyContractAndTransactoinStatus(mysql.DB, transaction.SourceID, status, contractAddress)

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
	}
	//
	//time.Sleep(30 * time.Second)
	//	}
	//}()
	return
}
