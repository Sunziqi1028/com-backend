package event

import (
	"ceres/pkg/config"
	"ceres/pkg/initialization/eth"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/qiniu/x/log"
	"math/big"
	"strings"
	"sync"
	"time"
)

func StartListen() {
	var count = 0
	for {
		log.Info("event.StartListen No:", count)
		waitGroup := &sync.WaitGroup{}
		waitGroup.Add(1)

		listenEvent := func() {
			defer waitGroup.Done()
			SubEvent()
		}

		go listenEvent()
		if count != 0 {
			if err := eth.Init(); err != nil {
				log.Warn(err)
			}
		}

		log.Info("event.StartListen Wait start")
		waitGroup.Wait()
		log.Info("event.StartListen Wait over")

		eth.Close()

		log.Info("event.StartListen Sleep:", 5*time.Second)
		time.Sleep(5 * time.Second)

		count++
	}
}

func SubEvent() {
	log.Info("SubEvent enter")
	startupAbi := GetABI(StartupContract.Abi)
	log.Info("SubEvent GetABI Done")
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
				return
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
