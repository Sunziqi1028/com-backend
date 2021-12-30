package startup

import (
	model "ceres/pkg/model/startup"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/startup"
	"strconv"

	"github.com/qiniu/x/log"
)

// ListStartups get startup list
func ListStartups(ctx *router.Context) {
	var request model.ListStartupRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	var response model.ListStartupsResponse
	if err := service.ListStartups(0, &request, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// ListStartupsMe get my startup list
func ListStartupsMe(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var request model.ListStartupRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
		ctx.HandleError(err)
		return
	}

	var response model.ListStartupsResponse
	if err := service.ListStartups(comerID, &request, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// GetStartup get startup
func GetStartup(ctx *router.Context) {
	startupID, err := strconv.ParseUint(ctx.Param("startupID"), 0, 64)
	if err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid startup ID")
		ctx.HandleError(err)
		return
	}
	var response model.GetStartupResponse
	if err := service.GetStartup(startupID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}
