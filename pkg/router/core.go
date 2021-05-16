package router

import "github.com/gin-gonic/gin"

/// Response http base response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

/// ceres Context of HTTP
type Context struct {
	*gin.Context
}

/// ceres handler function
type HandlerFunc func(c *Context)

/// Wrap the gin router function with the ceres core context
func Wrap(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

/// OK return with 200
func (ctx *Context) OK(data interface{}) {
	ctx.JSON(
		200,
		data,
	)
}

/// ERROR with message 
func (ctx *Context) ERROR(code int, message string) {
	ctx.Context.JSON(
		200,
		&Response{
			Code:    code,
			Message: message,
			Data:    nil,
		},
	)
}

/// JSON return data with code
func (ctx *Context) JSON(code int, data interface{}) {
	ctx.Context.JSON(
		200,
		&Response{
			Code:    code,
			Message: "Success",
			Data:    data,
		},
	)
}
