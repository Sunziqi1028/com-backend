package account

import (
	"ceres/pkg/model/account"
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	service "ceres/pkg/service/account"
	"ceres/pkg/utility/auth"
	"ceres/pkg/utility/jwt"
	"strings"

	"github.com/qiniu/x/log"
)

// LoginWithGithubCallback login with github oauth
func LoginWithGithubCallback(ctx *router.Context) {
	code := ctx.Query("code")
	if code == "" {
		err := router.ErrBadRequest.WithMsg("Code missed")
		ctx.HandleError(err)
		return
	}
	client := auth.NewGithubOauthClient(code)
	var (
		comerID uint64
		err     error
	)
	header := ctx.GetHeader("X-COMUNION-AUTHORIZATION")
	if strings.Trim(header, " ") == "" {
		comerID = 0
	} else {
		comerID, err = jwt.Verify(header)
	}
	if err != nil || comerID == 0 {
		var response account.ComerLoginResponse
		if err = service.LoginWithOauth(client, model.GithubOauth, &response); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(response)
	} else {
		if err = service.LinkOauthAccountToComer(comerID, client, model.GithubOauth); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(nil)
	}
}

// LoginWithGoogleCallback login with google oauth callback
func LoginWithGoogleCallback(ctx *router.Context) {
	code := ctx.Query("code")
	if code == "" {
		err := router.ErrBadRequest.WithMsg("Code missed")
		ctx.HandleError(err)
		return
	}
	client := auth.NewGoogleClient(code)
	var (
		comerID uint64
		err     error
	)
	comunioAuthHeader := ctx.GetHeader("X-COMUNION-AUTHORIZATION")
	if strings.Trim(comunioAuthHeader, " ") == "" {
		comerID = 0
		err = nil
	} else {
		comerID, err = jwt.Verify(comunioAuthHeader)
	}

	if err != nil || comerID == 0 {
		var response account.ComerLoginResponse
		if err = service.LoginWithOauth(client, model.GoogleOauth, &response); err != nil {
			ctx.HandleError(err)
			return
		}
		ctx.OK(response)
	} else {
		if err = service.LinkOauthAccountToComer(comerID, client, model.GoogleOauth); err != nil {
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
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Warn(err)
		err = router.ErrBadRequest.WithMsg(err.Error())
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
