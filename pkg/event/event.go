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
		log.Info("listen for contract:", config.Eth.StartupContractAddress)
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
				switch vLog.Topics[0].Hex() {
				case StartupContract.EventHex:
					intr, err := startupAbi.Events[StartupContract.Event].Inputs.UnpackValues(vLog.Data)
					if err != nil {
						log.Info(err)
						continue
					}
					go HandleStartup(intr[2].(common.Address).String(), intr[1], vLog.TxHash.String())
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
