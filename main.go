// @EgoctlOverwrite NO
// @EgoctlGenerateTime 20210223_202936
package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egovernor"
	"ceres/pkg/invoker"
	"ceres/pkg/router"
)

func main() {
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
