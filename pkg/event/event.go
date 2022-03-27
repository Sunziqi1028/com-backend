package event

import (
	"ceres/pkg/config"
	"ceres/pkg/initialization/eth"
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	model "ceres/pkg/model/startup"
	"context"
	"errors"
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
				log.Fatal(err)
			case vLog := <-logs:
				switch vLog.Topics[0].Hex() {
				case StartupContract.EventHex:
					intr, err := startupAbi.Events[StartupContract.Event].Inputs.UnpackValues(vLog.Data)
					if err != nil {
						log.Info(err)
						continue
					}

					startupTemp := intr[1].(struct {
						Name          string         `json:"name"`
						Mode          uint8          `json:"mode"`
						Hashtag       []string       `json:"hashtag"`
						Logo          string         `json:"logo"`
						Mission       string         `json:"mission"`
						TokenContract common.Address `json:"tokenContract"`
						Wallets       []struct {
							Name          string         `json:"name"`
							WalletAddress common.Address `json:"walletAddress"`
						} `json:"wallets"`
						Overview   string `json:"overview"`
						IsValidate bool   `json:"isValidate"`
					})

					comer := account.Comer{}
					comerAddress := intr[2].(common.Address).String()
					if err = account.GetComerByAddress(mysql.DB, comerAddress, &comer); err != nil {
						log.Warn(err)
						continue
					}

					if comer.ID == 0 {
						log.Warn(errors.New("comer does not exit"))
						continue
					}

					startup := model.Startup{
						ComerID:              comer.ID,
						Name:                 startupTemp.Name,
						Mode:                 model.Mode(startupTemp.Mode),
						Logo:                 startupTemp.Logo,
						Mission:              startupTemp.Mission,
						TokenContractAddress: startupTemp.TokenContract.String(),
						Overview:             startupTemp.Overview,
					}

					if err := model.CreateStartup(mysql.DB, &startup); err != nil {
						log.Warn(err)
						continue
					}
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
