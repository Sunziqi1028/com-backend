package account

import (
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"
	"strconv"
)

// ListAccounts list all accounts of the Comer
func ListAccounts(ctx *router.Context) {
	uin, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	response, err := service.GetComerAccounts(uin)
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return
	}
	ctx.OK(response)
}

// UnlinkAccount unlink accounts for the Comer
func UnlinkAccount(ctx *router.Context) {
	uin, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	identifier, err := strconv.ParseInt(ctx.Query("identifer"), 10, 64)
	if err != nil {
		ctx.ERROR(router.ErrParametersInvaild, err.Error())
		return
	}
	err = service.UnlinkComerAccount(uin, uint64(identifier))
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return
	}
	ctx.OK(nil)
}

// LinkWithGithub link current account with github
func LinkWithGithub(_ *router.Context) {

}

// LinkWithGithub link current account with github
func LinkWithMetamask(_ *router.Context) {

}
