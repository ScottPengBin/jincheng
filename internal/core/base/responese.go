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

type PaginateData struct {
	Records  interface{} `json:"records"`
	PageNum  int         `json:"page_num"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
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

// Success 成功
func (r *Response) Success(data interface{}) {
	r.returnJson(
		code,
		code,
		"成功",
		"",
		data,
	)
}

// Paginate 分页
func (r *Response) Paginate(data interface{}, total int64, param ReqPaginateParam) {
	var pd PaginateData
	pd.Records = data
	pd.Total = total
	pd.PageNum = param.GetPageNum()
	pd.PageSize = param.GetPageSize()
	r.Success(pd)
}

//ErrorParam 错误
func (r *Response) ErrorParam(err string) {
	r.ctx.JSON(http2.StatusOK, &Result{
		Code:    10010,
		Err:     err,
		Success: false,
	})
}

//Unauthorized 未授权登录
func (r *Response) Unauthorized(err string) {
	r.ctx.JSON(http2.StatusUnauthorized, &Result{
		Code:    10010,
		Err:     err,
		Success: false,
	})
}
