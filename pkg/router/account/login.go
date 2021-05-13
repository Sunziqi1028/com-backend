package account

import (
	"ceres/pkg/utility/auth"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"
)

func LoginWithGithub(_ *gin.Context) {

}

func LoginWithFacebook(_ *gin.Context) {

}

func LoginWithTwitter(_ *gin.Context) {

}

func LoginWithLinkedIn(_ *gin.Context) {

}

/// loginWithOauth is the abstraction of the standard login flow
func loginWithOauth(context *gin.Context, client auth.OauthClient) {
	// firsly auth the login
	// next to get the user profile
	// do the service logic
	_, err := client.GetUserProfile()
	if err != nil {
		elog.Error("Login with oauth faild")
		context.JSON(200, nil)
	}

	// use the account object to handle the logic if comer is exits then sign a new JWT Token 
}

func LoginWithMetamask(_ *gin.Context) {

}

func LoginWithIamtoken(_ *gin.Context) {

}
