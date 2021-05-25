package middleware

import (
	"ceres/pkg/utility/jwt"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"
)

const (
	AuthorizationHeader      = "X-COMUNION-AUTHORIZATION"
	AuthorizationErrorHeader = "X-COMUNION-AUTHFAILED"
	ComerUinContextKey       = "COMUNIONCOMERUIN"
	ComerRoleContextKey      = "COMUNIONROLE"
	ComerGuestRole           = "Guest"
	ComerLoginedRole         = "Comer"
)

/// JwtAuthorizationMiddleware
/// handle the authorization
func JwtAuthorizationMiddleware(ctx *gin.Context) {

	token := ctx.Request.Header[AuthorizationHeader]
	if token == nil || len(token) == 0 {
		ctx.Keys[ComerUinContextKey] = 0
		ctx.Keys[ComerRoleContextKey] = ComerGuestRole
		ctx.Next()
	} else {
		uin, err := jwt.Verify(token[0])
		if err != nil {
			elog.Warnf("Verify the request token failed %v", err)
			ctx.Next()
			ctx.Header(AuthorizationErrorHeader, err.Error()) // return the error to the client
		} else {
			ctx.Keys[ComerUinContextKey] = uin
			ctx.Keys[ComerRoleContextKey] = ComerLoginedRole
			ctx.Next()
		}
	}
}
