package http

import (
	"ceres/pkg/router"
	"ceres/pkg/router/account"
	"ceres/pkg/router/bounty"
	"ceres/pkg/router/image"
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

		oauthLogin.Use(middleware.ComerAuthorizationMiddleware())
		oauthLogin.POST("/register", router.Wrap(account.RegisterWithOauth))
	}

	// web3 login router
	web3Login := Gin.Group("/account/eth")
	{
		web3Login.Use(middleware.GuestAuthorizationMiddleware())
		web3Login.GET("/nonce", router.Wrap(account.GetBlockchainLoginNonce))
		web3Login.POST("/wallet/login", router.Wrap(account.LoginWithWallet))
	}

	// accounts operation router
	accountPriv := Gin.Group("/account")
	{
		accountPriv.Use(middleware.ComerAuthorizationMiddleware())
		// basic operations
		accountPriv.GET("/list", router.Wrap(account.ListAccounts))
		accountPriv.GET("/user/info", router.Wrap(account.UserInfo))
		accountPriv.POST("/eth/wallet/link", router.Wrap(account.LinkWithWallet))
		accountPriv.DELETE("/:accountID/unlink", router.Wrap(account.UnlinkAccount))
		// profile operations
		accountPriv.GET("/profile", router.Wrap(account.GetProfile))
		accountPriv.POST("/profile", router.Wrap(account.CreateProfile))
		accountPriv.PUT("/profile", router.Wrap(account.UpdateProfile))

		// comer operations
		accountPriv.POST("/comer/:comerID/follow", router.Wrap(account.FollowComer))
		accountPriv.DELETE("/comer/:comerID/unfollow", router.Wrap(account.UnfollowComer))
		accountPriv.GET("/comer/:comerID/followedByMe", router.Wrap(account.ComerFollowedByMe))
	}

	// accounts operation router
	accountsPub := Gin.Group("/account")
	{
		accountsPub.Use(middleware.GuestAuthorizationMiddleware())
		accountsPub.GET("/comer/:comerID", router.Wrap(account.GetComerInfo))
		accountsPub.GET("/comer/address/:address", router.Wrap(account.GetComerInfoByAddress))
	}

	coresPriv := Gin.Group("/cores")
	{
		coresPriv.Use(middleware.ComerAuthorizationMiddleware())
		coresPriv.GET("/startups/me", router.Wrap(startup.ListStartupsMe))
		coresPriv.POST("/startups/:startupID/follow", router.Wrap(startup.FollowStartup))
		coresPriv.DELETE("/startups/:startupID/unfollow", router.Wrap(startup.UnfollowStartup))
		coresPriv.GET("/startups/follow", router.Wrap(startup.ListFollowStartups))
		coresPriv.GET("/startups/participate", router.Wrap(startup.ListParticipateStartups))
		coresPriv.GET("/startups/:startupID/teamMembers", router.Wrap(startup.ListStartupTeamMembers))
		coresPriv.POST("/startups/:startupID/teamMembers/:comerID", router.Wrap(startup.CreateStartupTeamMember))
		coresPriv.PUT("/startups/:startupID/teamMembers/:comerID", router.Wrap(startup.UpdateStartupTeamMember))
		coresPriv.DELETE("/startups/:startupID/teamMembers/:comerID", router.Wrap(startup.DeleteStartupTeamMember))
		coresPriv.PUT("/startups/:startupID/basicSetting", router.Wrap(startup.UpdateStartupBasicSetting))
		coresPriv.PUT("/startups/:startupID/financeSetting", router.Wrap(startup.UpdateStartupFinanceSetting))
		coresPriv.GET("/startups/:startupID/followedByMe", router.Wrap(startup.StartupFollowedByMe))
	}

	coresPub := Gin.Group("/cores")
	{
		coresPub.Use(middleware.GuestAuthorizationMiddleware())
		coresPub.GET("/startups", router.Wrap(startup.ListStartups))
		coresPub.GET("/startups/:startupID", router.Wrap(startup.GetStartup))
		coresPub.GET("/startups/name/:name/isExist", router.Wrap(startup.StartupNameIsExist))
		coresPub.GET("/startups/tokenContract/:tokenContract/isExist", router.Wrap(startup.StartupTokenContractIsExist))
		coresPub.GET("/startups/member/:comerID", router.Wrap(startup.ListBeMemberStartups))
		//coresPub.GET("/startups/:startupId/setting", router.Wrap(startup.GetStartupSetting))
	}

	// misc operation router
	misc := Gin.Group("/misc")
	{
		misc.Use(middleware.ComerAuthorizationMiddleware())
		misc.POST("/upload", router.Wrap(upload.Upload))
	}

	// meta information
	meta := Gin.Group("/meta")
	{
		meta.Use(middleware.GuestAuthorizationMiddleware())
		meta.GET("/tags", router.Wrap(tag.GetTagList))
		meta.GET("/images", router.Wrap(image.GetImageList))
	}
	bounties := Gin.Group("/bounty")
	{
		bounties.Use(middleware.GuestAuthorizationMiddleware())
		bounties.GET("/startups/:comerID", router.Wrap(bounty.GetComerStartups))
		bounties.POST("/detail", router.Wrap(bounty.CreateBounty))
	}
	return
}
