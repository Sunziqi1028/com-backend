// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210223_202936
// @Deprecated： nouse api just an example
package router

import (
	"ceres/pkg/router/api"
	"ceres/pkg/router/core"

	"github.com/gin-gonic/gin"
)

func InitEnum(r gin.IRoutes) {
	r.GET("/api/enums/:id", core.Handle(api.EnumInfo))
	r.GET("/api/enums", core.Handle(api.EnumList))
	r.POST("/api/enums", core.Handle(api.EnumCreate))
	r.PUT("/api/enums/:id", core.Handle(api.EnumUpdate))
	r.DELETE("/api/enums/:id", core.Handle(api.EnumDelete))
}
