package account

import (
	"ceres/pkg/initialization/redis"
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"
	"context"
	"fmt"
	"strconv"

	"github.com/gotomicro/ego/core/elog"
)

// ListAccounts list all accounts of the Comer
func ListAccounts(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	fmt.Println(comerID)
	response, err := service.GetComerAccounts(comerID)
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return
	}

	ctx.OK(response)
}

// UnlinkAccount unlink accounts for the Comer
func UnlinkAccount(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	accountID, err := strconv.ParseUint(ctx.Param("accountID"), 0, 64)
	if err != nil {
		ctx.ERROR(router.ErrParametersInvaild, err.Error())
		return
	}
	err = service.UnlinkComerAccount(comerID, accountID)
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return
	}

	ctx.OK(nil)
}

// LinkWithWallet link current account with wallet
func LinkWithWallet(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	signature := &model.EthSignatureObject{}
	err := ctx.BindJSON(signature)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"wrong metamask login parameter",
		)
		return
	}

	//get nonce
	nonce, err := redis.Client.Get(context.TODO(), signature.Address)
	if err != nil {
		elog.Errorf("Comunion redis get key failed %v", err)
		return
	}

	err = service.LinkEthAccountToComer(
		comerID,
		signature.Address,
		signature.Signature,
		nonce,
	)

	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}

	ctx.OK(nil)
}

// CheckComerExists
func CheckComerExists(ctx *router.Context) {
	oin := ctx.Query("oin")
	result, err := service.CheckComerExists(oin)
	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}

	ctx.OK(result)
}
