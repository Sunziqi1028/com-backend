package web3

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gotomicro/ego/core/elog"
)

const (
	startupAddress string = "0x865F46B2aF27f76D41FaBE8dE2495911E487c8cE"
)

// Init the redis client
func Init() (err error) {
	go func() {
		client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
		if err != nil {
			panic(err)
		}
		query := ethereum.FilterQuery{
			Addresses: []common.Address{
				common.HexToAddress(startupAddress),
			},
		}
		logs := make(chan types.Log)
		sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			panic(err)
		}
		for {
			select {
			case err := <-sub.Err():
				elog.Fatal(err.Error())
			case vLog := <-logs:
				fmt.Println(vLog) // pointer to event log
			}
		}
		return
	}()
	return nil
}
