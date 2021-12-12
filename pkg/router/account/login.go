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
	"fmt"
	"net/http"
)

// LoginWithGithub login with github oauth
func LoginWithGithub(ctx *router.Context) {
	url := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%v&redirect_uri=%v", config.Github.ClientID, config.Github.CallbackURL)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

// LoginWithGithubCallback login with github oauth
func LoginWithGithubCallback(ctx *router.Context) {
	code := ctx.Query("code")
	if code == "" {
		ctx.ERROR(router.ErrParametersInvaild, "code missed")
		return
	}
	client := auth.NewGithubOauthClient(code)
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	if comerID == 0 {
		response, err := service.LoginWithOauth(client, model.GithubOauth)
		if err != nil {
			ctx.ERROR(router.ErrBuisnessError, err.Error())
			return
		}
		ctx.OK(response)
	} else {
		err := service.LinkOauthAccountToComer(comerID, client, model.GithubOauth)
		if err != nil {
			ctx.ERROR(router.ErrBuisnessError, err.Error())
			return
		}
		ctx.OK(nil)
	}
}

// LoginWithGoogle login with google oauth
func LoginWithGoogle(ctx *router.Context) {
	client := auth.NewGoogleClient("", "")
	url := client.AuthCodeURL(client.OauthState)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

// LoginWithGoogleCallback login with google oauth callback
func LoginWithGoogleCallback(ctx *router.Context) {
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
	client := auth.NewGoogleClient(state, code)
	comerID, _ := ctx.Keys[middleware.ComerUinContextKey].(uint64)
	if comerID == 0 {
		response, err := service.LoginWithOauth(client, model.GoogleOauth)
		if err != nil {
			ctx.ERROR(router.ErrBuisnessError, err.Error())
			return
		}
		ctx.OK(response)
	} else {
		err := service.LinkOauthAccountToComer(comerID, client, model.GoogleOauth)
		if err != nil {
			ctx.ERROR(router.ErrBuisnessError, err.Error())
			return
		}
		ctx.OK(nil)
	}
}

// GetBlockchainLoginNonce get the blockchain login nonce.
func GetBlockchainLoginNonce(ctx *router.Context) {
	address := ctx.Query("address")
	if address == "" {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"no web3 public key",
		)
		return
	}
	nonce, err := service.GenerateWeb3LoginNonce(address)
	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}
	ctx.OK(nonce)
}

// LoginWithWallet login with the wallet signature.
func LoginWithWallet(ctx *router.Context) {
	signature := &model.EthSignatureObject{}
	err := ctx.BindJSON(signature)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"wrong wallet login parameter",
		)
		return
	}

	nonce, err := redis.Client.Get(context.TODO(), signature.Address)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"wrong wallet login parameter",
		)
		return
	}
	// Replace the 0x prefix
	response, err := service.LoginWithEthWallet(
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
	ctx.OK(response)
}
