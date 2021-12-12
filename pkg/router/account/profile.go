package account

import (
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"
)

// CreateProfile create the profile
func CreateProfile(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	request := &model.CreateProfileRequest{}
	err := ctx.BindJSON(request)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"wrong metamask login parameter",
		)
		return
	}
	err = service.CreateComerProfile(comerID, request)
	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}

	ctx.OK(nil)
}

// GetProfile get current Comer profile
func GetProfile(ctx *router.Context) {
	uin, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	response, err := service.GetComerProfile(uin)
	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			"wrong metamask login parameter",
		)
		return
	}

	ctx.OK(response)
}

// UpdateProfile update the profile
func UpdateProfile(ctx *router.Context) {
	uin, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	request := &model.UpdateProfileRequest{}
	err := ctx.BindJSON(request)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"wrong metamask login parameter",
		)
		return
	}
	err = service.UpdateComerProfile(uin, request)
	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			"wrong metamask login parameter",
		)
		return
	}

	ctx.OK(nil)
}
