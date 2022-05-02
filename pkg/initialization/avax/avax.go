package avax

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/pubsub"
	"github.com/gorilla/websocket"
	"github.com/qiniu/x/log"
)

func Init() (err error) {
	dialer := websocket.Dialer{
		NetDial: func(netw, addr string) (net.Conn, error) {
			return net.Dial(netw, addr)
		},
	}

	httpHeader := http.Header{}
	log.Info("------------avax.Dial: ws://api.avax.network/ext/bc/X/events")
	conn, _, err := dialer.Dial("ws://api.avax.network/ext/bc/X/events", httpHeader)
	if err != nil {
		log.Warn("------------", err)
		log.Info("------------avax.Dial: ws://api.avax.network:9650/ext/bc/X/events")
		err = nil
		conn2, _, err2 := dialer.Dial("ws://api.avax.network:9650/ext/bc/X/events", httpHeader)
		if err2 != nil {
			log.Warn("------------", err2)
			return
		}
		conn = conn2
	}

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	readMsg := func() {
		defer waitGroup.Done()

		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				log.Info("------------", err)
				return
			}
			log.Info("------------messageType:", mt, "   msg:", string(msg))
			switch mt {
			case websocket.TextMessage:
				log.Info("------------", string(msg))
			default:
				log.Info("------------", mt, string(msg))
			}
		}
	}

	log.Info("------------before readMsg()")
	go readMsg()
	log.Info("------------after readMsg()")

	cmd := &pubsub.Command{NewSet: &pubsub.NewSet{}}
	cmdmsg, err := json.Marshal(cmd)
	if err != nil {
		log.Warn("------------", err)
		return
	}
	log.Info("------------cmdmsg:", string(cmdmsg))
	err = conn.WriteMessage(websocket.TextMessage, cmdmsg)
	if err != nil {
		log.Warn("------------", err)
		return
	}

	var addresses []string
	addresses = append(addresses, "0x7E94572BCc67B6eDa93DBa0493b681dC0ae9E964")
	cmd = &pubsub.Command{AddAddresses: &pubsub.AddAddresses{JSONAddresses: api.JSONAddresses{Addresses: addresses}}}
	cmdmsg, err = json.Marshal(cmd)
	if err != nil {
		log.Warn(err)
		return
	}
	log.Info("------------address cmdmsg:", string(cmdmsg))

	err = conn.WriteMessage(websocket.TextMessage, cmdmsg)
	if err != nil {
		log.Warn("------------", err)
		return
	}

	waitGroup.Wait()
	return
}
