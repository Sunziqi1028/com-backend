package http

import (
	"ceres/pkg/router"
	"ceres/pkg/router/account"
	"ceres/pkg/router/tags"

	"github.com/gotomicro/ego/server/egin"
)

var Gin *egin.Component

func Init() (err error) {
	Gin = egin.Load("server.http").Build()
	// register the router at this

	// Login Routers
	Gin.GET("/account/oauth/login/github", router.Wrap(account.LoginWithGithub))
	Gin.GET("/account/oauth/login/twitter", router.Wrap(account.LoginWithTwitter))
	Gin.GET("/account/oauth/login/facebook", router.Wrap(account.LoginWithFacebook))

	// Gin.Use()

	/// Below routers need the JWT verification middleware

	// Profile Routers

	// Bounty Routers

	// Disco Routers

	// Goverance Routers

	// tag list no authorization need but need some limit in gateway
	Gin.GET("/tags/startup/list", router.Wrap(tags.GetStartupTagList))
	Gin.GET("/tags/skill/list", router.Wrap(tags.GetStartupTagList))

	return
}
