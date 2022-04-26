package event

import (
	"ceres/pkg/config"
	"ceres/pkg/initialization/eth"
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/qiniu/x/log"
)

func SubEvent() {
	startupAbi := GetABI(StartupContract.Abi)
	select {
	case <-eth.EthSubChanel:
		log.Info("start sub eth client", config.Eth.StartupContractAddress)
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(config.Eth.Epoch),
			Addresses: []common.Address{common.HexToAddress(config.Eth.StartupContractAddress)},
		}
		logs := make(chan types.Log)
		sub, err := eth.Client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			panic(err)
		}
		for {
			select {
			case err = <-sub.Err():
				log.Warn(err)
			case vLog := <-logs:
				log.Info("vLog.BlockHash:", vLog.BlockHash.Hex())
				log.Info("vLog.BlockNumber:", vLog.BlockNumber)
				log.Info("vLog.TxHash:", vLog.TxHash.Hex())
				log.Info("vLog.Topics[0].Hex():", vLog.Topics[0].Hex())
				log.Info("StartupContract.EventHex:", StartupContract.EventHex)
				log.Info("startupContract.Event:", StartupContract.Event)
				log.Info("startupAbi.Events:", startupAbi.Events)
				log.Info("startupAbi:", startupAbi)
				log.Info("vLog.Data:", vLog.Data)
				switch vLog.Topics[0].Hex() {
				case StartupContract.EventHex:
					intr, err := startupAbi.Events[StartupContract.Event].Inputs.UnpackValues(vLog.Data)
					if err != nil {
						log.Info(err)
						continue
					}
					log.Info("intr:", intr)
					log.Info("intr[1]:", intr[1])
					log.Info("intr[2]:", intr[2].(common.Address).String())
					go HandleStartup(intr[2].(common.Address).String(), intr[1])
				}
			}
		}
	}
}

func GetABI(abiJSON string) abi.ABI {
	wrapABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic(err)
	}
	return wrapABI
}
