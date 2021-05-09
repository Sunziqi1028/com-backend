// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210223_202936
package router

import (
	"ceres/pkg/invoker"

	"github.com/gotomicro/ego/server/egin"
)

func ServeHTTP() *egin.Component {
	r := invoker.Gin

	// should register all router rules in this file

	InitEnum(r)

	return r
}
