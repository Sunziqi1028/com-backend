package eth

import (
	"ceres/pkg/config"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qiniu/x/log"
)

var Client *ethclient.Client
var EthSubChanel = make(chan struct{})

// Init the eth clien
func Init() (err error) {
	Client, err = ethclient.Dial(config.Eth.EndPoint + "/" + config.Eth.InfuraKey)
	if err != nil {
		log.Warn(err)
		return err
	}
	EthSubChanel <- struct{}{}
	return
}
