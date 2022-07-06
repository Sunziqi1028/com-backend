package eth

import (
	"ceres/pkg/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qiniu/x/log"
)

var Client *ethclient.Client
var EthSubChanel = make(chan struct{})

// Init the eth client
func Init() (err error) {
	log.Info("eth.Init ethclient.Dial:", config.Eth.EndPoint+"/"+config.Eth.InfuraKey)
	//Client, err = ethclient.Dial(config.Eth.EndPoint + "/" + config.Eth.InfuraKey)
	Client, err = ethclient.Dial(config.Eth.EndPoint)
	if err != nil {
		log.Warn(err)
		return err
	}
	log.Info("eth.Init ethclient.Dial done")
	EthSubChanel <- struct{}{}
	log.Info("eth.Init EthSubChanel <- struct{}{}")
	return
}

func Close() {
	log.Info("eth.Close start")
	Client.Close()
	EthSubChanel = make(chan struct{})
	log.Info("eth.Close end")
}

//func GetAllContractAddresses() {
//	ticker := time.NewTicker(10 * time.Minute)
//	go func() {
//		for {
//			t := ticker.C
//			fmt.Println("time now is :", t)
//			transactions, err := modelTransaction.GetTransaction(mysql.DB)
//			if err != nil {
//				return
//			}
//			for _, transaction := range transactions {
//				var contractChan = make(chan *modelBonuty.ContractInfoResponse, 1)
//				contractAddress, status := serviceTransaction.GetContractAddress(transaction.ChainID, transaction.TxHash)
//				contractInfo := &modelBonuty.ContractInfoResponse{
//					ContractAddress: contractAddress,
//					Status:          status,
//				}
//				select {
//				case contractChan <- contractInfo:
//					for contract := range contractChan {
//						serviceTransaction.UpdateBountyContractAndTransactoinStatus(mysql.DB, transaction.SourceID, contract.Status, contract.ContractAddress)
//						return
//					}
//				case <-time.After(5 * time.Second):
//					fmt.Println("get contract address time over!")
//				}
//				return
//			}
//		}
//	}()
//}
