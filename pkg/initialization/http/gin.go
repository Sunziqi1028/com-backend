package http

import (
	"ceres/pkg/router"
	"ceres/pkg/router/account"
	"ceres/pkg/router/bounty"
	"ceres/pkg/router/crowdfunding"
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
		coresPriv.GET("/startups/crowdfundable", router.Wrap(crowdfunding.SelectNonFundingStartups))
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
		// bounty
		coresPriv.GET("/bounties", router.Wrap(bounty.GetPublicBountyList))
		coresPriv.GET("/bounties/startup/:startupId", router.Wrap(bounty.GetBountyListByStartup))
		coresPriv.GET("/bounties/me/participated", router.Wrap(bounty.GetMyParticipatedBountyList))
		coresPriv.GET("/bounties/me/posted", router.Wrap(bounty.GetMyPostedBountyList))
		// crowdfunding

	}

	coresPub := Gin.Group("/cores")
	{
		coresPub.Use(middleware.GuestAuthorizationMiddleware())
		coresPub.GET("/startups", router.Wrap(startup.ListStartups))
		coresPub.GET("/startups/:startupID", router.Wrap(startup.GetStartup))
		coresPub.GET("/startups/name/:name/isExist", router.Wrap(startup.StartupNameIsExist))
		coresPub.GET("/startups/tokenContract/:tokenContract/isExist", router.Wrap(startup.StartupTokenContractIsExist))
		coresPub.GET("/startups/member/:comerID", router.Wrap(startup.ListBeMemberStartups))
		coresPub.GET("/startups/comer/:comerID", router.Wrap(startup.ListStartupsComer))
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
		// bounties.Use(middleware.GuestAuthorizationMiddleware())
		bounties.Use(middleware.ComerAuthorizationMiddleware())
		//bounties.GET("/startups/:comerID", router.Wrap(bounty.GetComerStartups))
		bounties.POST("/detail", router.Wrap(bounty.CreateBounty))
		bounties.GET("/detail/:bountyID", router.Wrap(bounty.GetBountyDetailByID))
		bounties.GET("/:bountyID/payment", router.Wrap(bounty.GetPaymentByBountyID))
		bounties.PUT("/:bountyID/close", router.Wrap(bounty.UpdateBountyStatus))
		bounties.POST("/add/deposit", router.Wrap(bounty.AddDeposit))
		bounties.PUT("/:bountyID/paid/status", router.Wrap(bounty.UpdatePaidStatusByBountyID))
		bounties.POST("/activities", router.Wrap(bounty.CreateActivities))
		bounties.POST("/applicants", router.Wrap(bounty.CreateApplicants))
		bounties.GET("/list/:bountyID/activities", router.Wrap(bounty.GetActivitiesLists))
		bounties.GET("/list/:bountyID/applicants", router.Wrap(bounty.GetAllApplicantsByBountyID))
		bounties.GET("/:bountyID/founder", router.Wrap(bounty.GetFounderByBountyID))
		bounties.GET("/:bountyID/approved", router.Wrap(bounty.GetApprovedApplicantByBountyID))
		bounties.GET("/:bountyID/deposit-records", router.Wrap(bounty.GetDepositRecords))
		bounties.PUT("/founder/:bountyID/approve", router.Wrap(bounty.UpdateFounderApprovedApplicant))
		bounties.PUT("/founder/:bountyID/unapprove", router.Wrap(bounty.UpdateFounderApprovedApplicant))
		bounties.GET("/:bountyID/startup", router.Wrap(bounty.GetStartupByBountyID))
		//bounties.GET("/:bountyID/role/founder", router.Wrap(bounty.GetStartupByBountyID))
		bounties.GET("/:bountyID/role/applicant", router.Wrap(bounty.GetBountyRoleByComerID))
		//bounties.PUT("/:bountyID/applicant/unlock", router.Wrap(bounty.UpdateFounderApprovedApplicant))
		//bounties.PUT("/:bountyID/founder/release", router.Wrap(bounty.UpdateFounderApprovedApplicant))

	}

	crowdfundingRouters := Gin.Group("/crowdfunding")
	{
		crowdfundingRouters.Use(middleware.ComerAuthorizationMiddleware())
		// crowdFunding
		crowdfundingRouters.POST("/", router.Wrap(crowdfunding.CreateCrowdFunding))

	}

	//testR := Gin.Group("/test")
	//{
	//	testR.Use(middleware.GuestAuthorizationMiddleware())
	//	testR.GET("/token/:comerId", router.Wrap(my.Token))
	//}
	return
}
