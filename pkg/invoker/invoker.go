// @EgoctlOverwrite NO
// @EgoctlGenerateTime 20210223_202936
// @Deprecated： replace to the initialization package
package invoker

import (
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egin"
)

var (
	Logger       *elog.Component
	Db           *egorm.Component
	Gin          *egin.Component
)

func Init() error {
	Logger = elog.DefaultLogger
	Gin = egin.Load("server.http").Build()
	Db = egorm.Load("mysql.user").Build()
	return nil
}