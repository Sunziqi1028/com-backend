package metrics

import "github.com/gotomicro/ego/server/egovernor"

var Vernor *egovernor.Component

func Init() (err error) {

	Vernor = egovernor.Load("server.governor").Build()

	return
}
