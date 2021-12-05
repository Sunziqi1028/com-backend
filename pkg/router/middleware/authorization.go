package middleware

import (
	"ceres/pkg/router"
	"ceres/pkg/utility/jwt"
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"
	"net/http"
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
		ctx.Keys = make(map[string]interface{})
		ctx.Keys[ComerUinContextKey] = 0
		ctx.Keys[ComerRoleContextKey] = ComerGuestRole
		ctx.Next()
	}
}

// JwtAuthorizationMiddleware  handle the authorization
func JwtAuthorizationMiddleware(ctx *gin.Context) {
	token := ctx.Request.Header[http.CanonicalHeaderKey(AuthorizationHeader)]

	if len(token) == 0 {
		ctx.Header(AuthorizationErrorHeader, "Have to login") // return the error to the client
		ctx.AbortWithStatusJSON(router.ErrForbidden, nil)
	}

	uin, err := jwt.Verify(token[0])
	if err != nil {
		elog.Warnf("Verify the request token failed %v", err)
		ctx.Header(AuthorizationErrorHeader, err.Error()) // return the error to the client
		ctx.AbortWithStatusJSON(router.ErrForbidden, nil)
	}

	ctx.Keys = make(map[string]interface{})
	ctx.Keys[ComerUinContextKey] = uin
	ctx.Keys[ComerRoleContextKey] = ComerLoginedRole
	ctx.Next()
}

//Jwt