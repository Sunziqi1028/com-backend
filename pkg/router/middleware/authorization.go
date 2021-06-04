package middleware

import (
	"ceres/pkg/utility/jwt"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"
)

// Middleware constraints
const (
	AuthorizationHeader      = "X-COMUNION-AUTHORIZATION"
	AuthorizationErrorHeader = "X-COMUNION-AUTHFAILED"
	ComerUinContextKey       = "COMUNIONCOMERUIN"
	ComerRoleContextKey      = "COMUNIONROLE"
	ComerGuestRole           = "Guest"
	ComerLoginedRole         = "Comer"
)

// ComerAuthorizationMiddleware return the comer authorization middleware
func ComerAuthorizationMiddleware() gin.HandlerFunc {
	return JwtAuthorizationMiddleware
}

// GuestAuthorizationMiddleware return the guest authorization middleware
func GuestAuthorizationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Keys[ComerUinContextKey] = 0
		ctx.Keys[ComerRoleContextKey] = ComerGuestRole
		ctx.Next()
	}
}

// JwtAuthorizationMiddleware  handle the authorization
func JwtAuthorizationMiddleware(ctx *gin.Context) {
	token := ctx.Request.Header[AuthorizationHeader]
	if len(token) == 0 {
		ctx.Header(AuthorizationErrorHeader, "Have to login") // return the error to the client
		ctx.JSON(403, nil)
	} else {
		uin, err := jwt.Verify(token[0])
		if err != nil {
			elog.Warnf("Verify the request token failed %v", err)
			ctx.Header(AuthorizationErrorHeader, err.Error()) // return the error to the client
			ctx.JSON(403, nil)
		} else {
			ctx.Keys[ComerUinContextKey] = uin
			ctx.Keys[ComerRoleContextKey] = ComerLoginedRole
			ctx.Next()
		}
	}
}
