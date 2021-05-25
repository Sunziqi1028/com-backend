package account

import (
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	service "ceres/pkg/service/account"
	"ceres/pkg/utility/auth"
	"encoding/hex"
	"strings"
)

/// LoginWithGithub
/// login with github oauth
func LoginWithGithub(ctx *router.Context) {
	requestToken := ctx.Query("request_token")
	if requestToken == "" {
		ctx.ERROR(router.ErrParametersInvaild, "request_token missed")
		return
	}
	client := auth.NewGithubOauthClient(requestToken)
	response, err := service.LoginWithOauth(client, model.GithubOauth)
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return
	}
	ctx.OK(response)
}

/// LoginWithFacebook
/// login with facebook oauth
func LoginWithFacebook(ctx *router.Context) {
	requestToken := ctx.Query("request_token")
	if requestToken == "" {
		ctx.ERROR(router.ErrParametersInvaild, "request_token missed")
		return
	}
	client := auth.NewFacebookClient(requestToken)
	response, err := service.LoginWithOauth(client, model.GithubOauth)
	if err != nil {
		ctx.ERROR(router.ErrBuisnessError, err.Error())
		return
	}
	ctx.OK(response)
}

func LoginWithTwitter(_ *router.Context) {
	// TODO: should complete the twitter logic
}

func LoginWithLinkedIn(_ *router.Context) {
	// TODO: should complete the linkedin logic
}

/// GetBlockchainLoginNonce
/// get the blockchain login nonce
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

/// LoginWithMetamask
/// login with the metamask signature
func LoginWithMetamask(ctx *router.Context) {
	signature := &model.EthSignatureObject{}
	err := ctx.BindJSON(signature)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"wrong metamask login parameter",
		)
		return
	}

	// Replace the 0x prefix

	a := strings.Replace(signature.Address, "0x", "", 1)
	s := strings.Replace(signature.Signature, "0x", "", 1)
	m := strings.Replace(signature.MessageHash, "0x", "", 1)

	address, err := hex.DecodeString(a)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"illegal address",
		)
		return
	}
	sign, err := hex.DecodeString(s)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"illegal signature",
		)
	}
	message, err := hex.DecodeString(m)
	if err != nil {
		ctx.ERROR(
			router.ErrParametersInvaild,
			"illegal message",
		)
		return
	}

	resposne, err := service.VerifyEthSignatureAndLogin(
		address,
		message,
		sign,
		model.MetamaskEth,
	)

	if err != nil {
		ctx.ERROR(
			router.ErrBuisnessError,
			err.Error(),
		)
		return
	}
	ctx.OK(resposne)
}
