// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210223_200720
package router

import (
	"ceres/pkg/router/api"
    "github.com/gin-gonic/gin"
    "ceres/pkg/router/core"
)

func InitEnum(r gin.IRoutes) {
    r.GET("/api/enums/:id", core.Handle(api.EnumInfo))
    r.GET("/api/enums", core.Handle(api.EnumList))
    r.POST("/api/enums", core.Handle(api.EnumCreate))
    r.PUT("/api/enums/:id", core.Handle(api.EnumUpdate))
    r.DELETE("/api/enums/:id", core.Handle(api.EnumDelete))
}
