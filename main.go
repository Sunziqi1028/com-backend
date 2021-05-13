// @EgoctlOverwrite NO
// @EgoctlGenerateTime 20210223_202936
package main

import (
	"ceres/pkg/invoker"
	"ceres/pkg/router"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egovernor"
)

func main() {
	// Order:
	// init the logger
	// init the gorm
	// init the redis
	// init the utility
	// init the grpc
	// init the gin

	if err := ego.New().
		Invoker(invoker.Init).
		Serve(
			egovernor.Load("server.governor").Build(),
			router.ServeHTTP(),
		).
		Run(); err != nil {
		elog.Panic(err.Error())
	}
}
