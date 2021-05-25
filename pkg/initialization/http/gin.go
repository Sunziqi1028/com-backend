package http

import (
	"ceres/pkg/router"
	"ceres/pkg/router/account"

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

	return
}
