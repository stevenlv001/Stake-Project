package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

const (
	CodeSuccess      = 0
	CodeBadRequest   = 400
	CodeUnauthorized = 401
	CodeForbidden    = 403
	CodeNotFound     = 404
	CodeServerError  = 500
	CodeTooManyReq   = 429
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Data: data,
		Msg:  "success",
	})
}

func SuccessWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  msg,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
	})
}

func BadRequest(c *gin.Context, msg string) {
	Fail(c, CodeBadRequest, msg)
}

func Unauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: CodeUnauthorized,
		Msg:  msg,
	})
}

func ServerError(c *gin.Context, msg string) {
	Fail(c, CodeServerError, msg)
}

func TooManyRequests(c *gin.Context) {
	c.JSON(http.StatusTooManyRequests, Response{
		Code: CodeTooManyReq,
		Msg:  "请求过于频繁，请稍后再试",
	})
}
