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
	Client, err = ethclient.Dial("wss://api.avax-test.network/ext/bc/C/ws")
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
