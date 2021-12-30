package http

import (
	"ceres/pkg/router"
	"ceres/pkg/router/account"
	"ceres/pkg/router/middleware"
	"ceres/pkg/router/startup"
	"ceres/pkg/router/tag"
	"ceres/pkg/router/upload"

	"github.com/gotomicro/ego/server/egin"
)

// Gin instance
var Gin *egin.Component

// Init the Gin instance and the routers
func Init() (err error) {
	Gin = egin.Load("server.http").Build()
	// oauth login router
	oauthLogin := Gin.Group("/account/oauth")
	{
		oauthLogin.Use(middleware.GuestAuthorizationMiddleware())
		oauthLogin.GET("/github/login/callback", router.Wrap(account.LoginWithGithubCallback))
		oauthLogin.GET("/google/login/callback", router.Wrap(account.LoginWithGoogleCallback))
	}

	// web3 login router
	web3Login := Gin.Group("/account/eth")
	{
		web3Login.Use(middleware.GuestAuthorizationMiddleware())
		web3Login.GET("/nonce", router.Wrap(account.GetBlockchainLoginNonce))
		web3Login.POST("/wallet/login", router.Wrap(account.LoginWithWallet))
	}

	// accounts operation router
	accounts := Gin.Group("/account")
	{
		accounts.Use(middleware.ComerAuthorizationMiddleware())
		// basic operations
		accounts.GET("/list", router.Wrap(account.ListAccounts))
		accounts.POST("/eth/wallet/link", router.Wrap(account.LinkWithWallet))
		accounts.DELETE("/:accountID/unlink", router.Wrap(account.UnlinkAccount))
		// profile operations
		accounts.GET("/profile", router.Wrap(account.GetProfile))
		accounts.POST("/profile", router.Wrap(account.CreateProfile))
		accounts.PUT("/profile", router.Wrap(account.UpdateProfile))
	}

	coresPriv := Gin.Group("/cores")
	{
		coresPriv.Use(middleware.ComerAuthorizationMiddleware())
		coresPriv.GET("/startups/me", router.Wrap(startup.ListStartupsMe))
	}

	coresPub := Gin.Group("/cores")
	{
		coresPub.Use(middleware.GuestAuthorizationMiddleware())
		coresPub.GET("/startups", router.Wrap(startup.ListStartups))
		coresPub.GET("/startups/:startupID", router.Wrap(startup.GetStartup))
		//coresPub.GET("/startups/:startupId/setting", router.Wrap(startup.GetStartupSetting))
	}

	// misc operation router
	misc := Gin.Group("/misc")
	{
		misc.Use(middleware.ComerAuthorizationMiddleware())
		misc.GET("/upload/presign", router.Wrap(upload.GetPresignedURLForUpload))
	}

	// meta information
	meta := Gin.Group("/meta")
	{
		meta.Use(middleware.GuestAuthorizationMiddleware())
		meta.GET("/tags", router.Wrap(tag.GetTagList))
	}

	return
}
