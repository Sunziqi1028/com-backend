package middleware

import (
	"ceres/pkg/initialization/mysql"
	"ceres/pkg/model/account"
	model "ceres/pkg/model/account"
	"ceres/pkg/router"
	"ceres/pkg/utility/jwt"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/x/log"
	"gorm.io/gorm"
)

// Middleware constraints
const (
	AuthorizationHeader = "X-COMUNION-AUTHORIZATION"
	ComerUinContextKey  = "COMUNIONCOMERUIN"
	ComerRoleContextKey = "COMUNIONROLE"
	ComerGuestRole      = "Guest"
	ComerLoginedRole    = "Comer"
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
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, router.ErrUnauthorized)
		return
	}

	comerID, err := jwt.Verify(token[0])
	if err != nil {
		log.Warnf("Verify the request token failed %v", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, router.ErrUnauthorized.WithMsgf("Verify the request token failed %v", err.Error()))
		return
	}

	var comer account.Comer
	if err := model.GetComerByID(mysql.DB, comerID, &comer); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, router.ErrUnauthorized.WithMsg("comer does not exist"))
		}
		log.Warnf("get comer fail %v", err)
		return
	}

	ctx.Keys = make(map[string]interface{})
	ctx.Keys[ComerUinContextKey] = comerID
	ctx.Keys[ComerRoleContextKey] = ComerLoginedRole
	ctx.Next()
}

//Jwt
