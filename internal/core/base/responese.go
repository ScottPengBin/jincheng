package base

import (
	"github.com/gin-gonic/gin"
	http2 "net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Err     string      `json:"error"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
}

const code = 10000

type Response struct {
	ctx *gin.Context
}

func NewResponse(ctx *gin.Context) Response {
	return Response{
		ctx: ctx,
	}
}

func (r *Response) returnJson(code int, dataCode int, msg string, err string, data interface{}) {
	success := false
	if code == dataCode {
		success = true
	}

	r.ctx.JSON(http2.StatusOK, &Result{
		Code:    code,
		Data:    data,
		Msg:     msg,
		Success: success,
	})
}

func (r Response) Success(data interface{}) {
	r.returnJson(
		code,
		code,
		"成功",
		"",
		data,
	)
}
