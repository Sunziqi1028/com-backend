package account

import (
	"ceres/pkg/config"
	"ceres/pkg/initialization/redis"
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/router/middleware"
	service "ceres/pkg/service/account"
	"ceres/pkg/utility/auth"
	"context"
	"net/http"
	"strconv"

	"github.com/gotomicro/ego/core/elog"
)

// LinkWithGithub link current account with github
func LinkWithGithub(ctx *router.Context) {
	//uin, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	//requestToken := ctx.Query("request_token")
	//if requestToken == "" {
	//	ctx.ERROR(router.ErrParametersInvaild, "request_token missed")
	//	return
	//}
	//client := auth.NewGithubOauthClient(requestToken)
	//err := service.LinkOauthAccountToComer(uin, client, model.GithubOauth)
	//if err != nil {
	//	ctx.ERROR(router.ErrBuisnessError, err.Error())
	//	return
	//
	//}
	ctx.OK(nil)
}

// LinkWithGoogle link current account with google
func LinkWithGoogle(ctx *router.Context) {
	client := auth.NewGoogleClient(config.Google.LinkCallbackURL, "", "")
	url := client.AuthCodeURL(client.OauthState)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

// LinkWithGoogleCallback link current account with google callback
func LinkWithGoogleCallback(ctx *router.Context) {
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	state := ctx.Query("state")
	if state == "" {
		ctx.ERROR(router.ErrParametersInvaild, "state missed")
		return
	}
	code := ctx.Query("code")
	if code == "" {
		ctx.ERROR(router.ErrParametersInvaild, "code missed")
		return
	}
	client := auth.NewGoogleClient(config.Google.LinkCallbackURL, "", "")
	err := service.LinkOauthAccountToComer(comerID, client, model.GoogleOauth)
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return

	}
	ctx.OK(nil)
}

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
	identifier, err := strconv.ParseInt(ctx.Query("identifier"), 10, 64)
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
