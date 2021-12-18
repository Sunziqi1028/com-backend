package account

import (
	"ceres/pkg/config"
	"ceres/pkg/model/account"
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	service "ceres/pkg/service/account"
	"ceres/pkg/utility/auth"
	"ceres/pkg/utility/jwt"
	"fmt"
	"net/http"
)

// LoginWithGithub login with github oauth
func LoginWithGithub(ctx *router.Context) {
	url := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%v&redirect_uri=%v&state=%v", config.Github.ClientID, config.Github.CallbackURL, ctx.GetHeader("X-COMUNION-AUTHORIZATION"))
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

// LoginWithGithubCallback login with github oauth
func LoginWithGithubCallback(ctx *router.Context) {
	code := ctx.Query("code")
	if code == "" {
		err := router.ErrBadRequest.WithMsg("Code missed")
		ctx.HandleError(err)
		return
	}
	client := auth.NewGithubOauthClient(code)
	comerID, err := jwt.Verify(ctx.GetHeader(ctx.GetHeader("X-COMUNION-AUTHORIZATION")))
	if err != nil || comerID == 0 {
		var response account.ComerLoginResponse
		if err := service.LoginWithOauth(client, model.GithubOauth, &response); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(response)
	} else {
		if err := service.LinkOauthAccountToComer(comerID, client, model.GithubOauth); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(nil)
	}
}

// LoginWithGoogle login with google oauth
func LoginWithGoogle(ctx *router.Context) {
	jwtHeader := ctx.GetHeader("X-COMUNION-AUTHORIZATION")
	client := auth.NewGoogleClient(jwtHeader, "")
	url := client.AuthCodeURL(client.OauthState)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

// LoginWithGoogleCallback login with google oauth callback
func LoginWithGoogleCallback(ctx *router.Context) {
	code := ctx.Query("code")
	if code == "" {
		err := router.ErrBadRequest.WithMsg("Code missed")
		ctx.HandleError(err)
		return
	}
	client := auth.NewGoogleClient("", code)
	comerID, err := jwt.Verify(ctx.GetHeader(ctx.GetHeader("X-COMUNION-AUTHORIZATION")))
	if err != nil || comerID == 0 {
		var response account.ComerLoginResponse
		if err := service.LoginWithOauth(client, model.GoogleOauth, &response); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(response)
	} else {
		if err := service.LinkOauthAccountToComer(comerID, client, model.GoogleOauth); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(nil)
	}
}

// GetBlockchainLoginNonce get the blockchain login nonce.
func GetBlockchainLoginNonce(ctx *router.Context) {
	address := ctx.Query("address")
	if address == "" {
		err := router.ErrBadRequest.WithMsg("Invalid address")
		ctx.HandleError(err)
		return
	}

	var nonce account.WalletNonceResponse
	if err := service.GenerateWeb3LoginNonce(address, &nonce); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(nonce)
}

// LoginWithWallet login with the wallet signature.
func LoginWithWallet(ctx *router.Context) {
	var request model.EthLoginRequest
	if err := ctx.BindJSON(&request); err != nil {
		err = router.ErrBadRequest.WithMsg("Invalid data format")
		ctx.HandleError(err)
		return
	}

	var response account.ComerLoginResponse
	if err := service.LoginWithEthWallet(request.Address, request.Signature, &response); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.OK(response)
}
