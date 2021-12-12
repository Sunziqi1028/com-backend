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

	// account without token
	accountWhite := Gin.Group("/account/white")
	{
		accountWhite.GET("/check", router.Wrap(account.CheckComerExists))
	}

	// oauth login router
	oauthLogin := Gin.Group("/account/oauth")
	{
		oauthLogin.Use(middleware.GuestAuthorizationMiddleware())
		oauthLogin.GET("/github/login", router.Wrap(account.LoginWithGithub))
		oauthLogin.GET("/github/login/callback", router.Wrap(account.LoginWithGithubCallback))
		oauthLogin.GET("/google/login", router.Wrap(account.LoginWithGoogle))
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

	//coresPriv := Gin.Group("/cores")
	//{
	//	coresPriv.Use(middleware.ComerAuthorizationMiddleware())
	//	coresPriv.GET("/startups/me", router.Wrap(account.ListAccounts))
	//	coresPriv.GET("/startups/:startupId/bounties/me", router.Wrap(account.ListAccounts))
	//}

	// startups operation router
	//coresPub := Gin.Group("/cores")
	//{
	//	coresPub.Use(middleware.GuestAuthorizationMiddleware())
	//	// basic operations
	//	coresPub.GET("/startups", router.Wrap(account.ListAccounts))
	//	coresPub.GET("/startups/:startupId", router.Wrap(account.LinkWithGithub))
	//	coresPub.GET("/startups/:startupId/bounties/", router.Wrap(account.ListAccounts))
	//	coresPub.GET("/bounties/:bountyId", router.Wrap(account.ListAccounts))
	//	coresPub.GET("/bounties", router.Wrap(account.ListAccounts))
	//	coresPub.GET("/dios/:dioId", router.Wrap(account.ListAccounts))
	//	coresPub.GET("/startups/:startupId/dios", router.Wrap(account.ListAccounts))
	//	coresPub.GET("/startups/:startupId/dios/:dioId", router.Wrap(account.ListAccounts))
	//}

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
