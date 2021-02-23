// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210223_200720
package router

import (
	"github.com/gotomicro/ego/server/egin"
	"ceres/pkg/invoker"
)

func ServeHTTP() *egin.Component {
	r := invoker.Gin
    
    InitEnum(r)
    
	return r
}
