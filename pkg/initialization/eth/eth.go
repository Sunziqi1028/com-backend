package eth

import (
	"ceres/pkg/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qiniu/x/log"
)

var Client *ethclient.Client
var RPCClient *ethclient.Client
var EthSubChanel = make(chan struct{})

// Init the eth client
func Init() (err error) {
	//log.Info("eth.Init ethclient.Dial:", config.Eth.EndPoint+"/"+config.Eth.InfuraKey)
	//Client, err = ethclient.Dial(config.Eth.EndPoint + "/" + config.Eth.InfuraKey)
	// Client, err = ethclient.Dial("wss://api.avax-test.network/ext/bc/C/ws")
	log.Info("eth.Init ethclient_wss.Dial:", config.Eth.WSSEndPoint)
	Client, err = ethclient.Dial(config.Eth.WSSEndPoint)
	if err != nil {
		log.Warn(err)
		return err
	}
	log.Info("eth.Init ethclient_rpc.Dial:", config.Eth.RPCEndPoint)
	RPCClient, err = ethclient.Dial(config.Eth.RPCEndPoint)
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
