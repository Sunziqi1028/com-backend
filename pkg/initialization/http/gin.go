package http

import (
	"ceres/pkg/router"
	"ceres/pkg/router/account"
	"ceres/pkg/router/middleware"
	"ceres/pkg/router/tags"
	"ceres/pkg/router/upload"

	"github.com/gotomicro/ego/server/egin"
)

// Gin instance
var Gin *egin.Component

// Init the Gin instance and the routers
func Init() (err error) {
	Gin = egin.Load("server.http").Build()
	// register the router at this

	// oauth login router
	oauthLogin := Gin.Group("/account/oauth/login")
	{
		oauthLogin.Use(middleware.GuestAuthorizationMiddleware())
		oauthLogin.POST("/github", router.Wrap(account.LoginWithGithub))
		oauthLogin.POST("/facebook", router.Wrap(account.LoginWithFacebook))
	}

	// web3 login router
	web3Login := Gin.Group("/account/eth/login")
	{
		web3Login.Use(middleware.GuestAuthorizationMiddleware())
		web3Login.GET("/nonce", router.Wrap(account.GetBlockchainLoginNonce))
		web3Login.POST("/metamask", router.Wrap(account.LoginWithMetamask))
	}

	// accounts operation router
	accounts := Gin.Group("/account")
	{
		accounts.Use(middleware.ComerAuthorizationMiddleware())
		// basic operations
		accounts.GET("/list", router.Wrap(account.ListAccounts))
		accounts.POST("/oauth/link/gihub", router.Wrap(account.LinkWithGithub))
		accounts.POST("/eth/link/metamask", router.Wrap(account.LinkWithMetamask))
		accounts.DELETE("/unlink", router.Wrap(account.UnlinkAccount))
		// profile operations
		accounts.POST("/profile", router.Wrap(account.CreateProfile))
		accounts.GET("/profile", router.Wrap(account.GetProfile))
		accounts.PUT("/profile", router.Wrap(account.UpdateProfile))
	}

	// misc operation router
	misc := Gin.Group("/misc")
	{
		misc.Use(middleware.ComerAuthorizationMiddleware())
		misc.GET("/upload/presign", router.Wrap(upload.GetPresignedURLForUpload))
	}

	// meta informations
	meta := Gin.Group("/meta")
	{
		meta.Use(middleware.GuestAuthorizationMiddleware())
		meta.GET("/tag/startup", router.Wrap(tags.GetStartupTagList))
		meta.GET("/tag/skill", router.Wrap(tags.GetSkillTagList))
	}

	return
}
