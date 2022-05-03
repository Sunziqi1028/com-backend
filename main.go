// @EgoctlOverwrite NO
// @EgoctlGenerateTime 20210223_202936
package main

import (
	"ceres/pkg/event"
	"ceres/pkg/initialization/config"
	"ceres/pkg/initialization/eth"
	"ceres/pkg/initialization/http"
	"ceres/pkg/initialization/logger"
	"ceres/pkg/initialization/metrics"
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/initialization/redis"
	"ceres/pkg/initialization/s3"
	"ceres/pkg/initialization/utility"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
)

func main() {
	// Order
	// init the config file
	// init the config file
	// init the logger
	// init the gorm
	// init the redis
	// init the metrics
	// init the utility
	// init the grpc
	// init the gin
	// init the web3
	go event.StartListen()
	//go avax.Init()
	if err := ego.New().Invoker(
		config.Init,
		logger.Init,
		mysql.Init,
		redis.Init,
		metrics.Init,
		utility.Init,
		http.Init,
		s3.Init,
		eth.Init,
	).Serve(
		metrics.Vernor,
		http.Gin,
	).Run(); err != nil {
		elog.Panic(err.Error())
	}
}
