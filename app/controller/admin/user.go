package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	reqAdmin "jincheng/app/request/admin"
	"jincheng/internal/core/base"
	"jincheng/internal/service/admin"
	"strconv"
)

type UserController struct {
	service *admin.UserService
}

func NewUserController(service *admin.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

var UserProvider = wire.NewSet(NewUserController, admin.NewUserService)

//QueryUsersByPage 用户管理列表
func (c *UserController) QueryUsersByPage(ctx *gin.Context) {
	output := base.NewResponse(ctx)
	var param reqAdmin.UserSearchParam
	param.PageNum, _ = strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	param.PageSize, _ = strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))
	param.Account = ctx.Query("account")
	param.Name = ctx.Query("name")
	param.Mobile = ctx.Query("mobile")

	res, total := c.service.QueryUsersByPage(&param)

	output.Paginate(res, total, &param)
}
