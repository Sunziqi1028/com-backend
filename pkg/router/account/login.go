package account

import (
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	service "ceres/pkg/service/account"
	"ceres/pkg/utility/auth"
)

/// LoginWithGithub
/// login with github oauth
func LoginWithGithub(ctx *router.Context) {
	requestToken := ctx.Query("request_token")
	if requestToken == "" {
		ctx.ERROR(400, "request_token missed")
		return
	}
	client := auth.NewGithubOauthClient(requestToken)
	response, err := service.LoginWithOauth(client, model.GithubOauth)
	if err != nil {
		ctx.ERROR(500, err.Error())
		return
	}
	ctx.OK(response)
	return
}

func LoginWithFacebook(_ *router.Context) {

}

func LoginWithTwitter(_ *router.Context) {

}

func LoginWithLinkedIn(_ *router.Context) {

}

func LoginWithMetamask(_ *router.Context) {

}

func LoginWithImtoken(_ *router.Context) {

}
