package account

import (
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"

	"github.com/qiniu/x/log"
)

// CreateProfile create the profile
func CreateProfile(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	request := &model.CreateProfileRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		err = router.ErrBadRequest.WithMsg(err.Error())
		log.Warn(err)
		ctx.HandleError(err)
		return
	}

	if err := service.CreateComerProfile(comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}

// GetProfile get current Comer profile
func GetProfile(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	var response model.ComerProfileResponse
	if err := service.GetComerProfile(comerID, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}

// UpdateProfile update the profile
func UpdateProfile(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	request := &model.UpdateProfileRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		err = router.ErrBadRequest.WithMsg(err.Error())
		log.Warn(err)
		ctx.HandleError(err)
		return
	}

	if err := service.UpdateComerProfile(comerID, request); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nil)
}
