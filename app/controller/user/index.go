package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/internal/core/base"
	"jincheng/internal/service/user"
)

var Provider = wire.NewSet(NewController,user.NewService)

func NewController(s user.Service) Controller {
	return Controller{
		service: s,
	}
}

type Controller struct {
	service user.Service
}

type Param struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (c *Controller) GetList(ctx *gin.Context) {
	var param Param
	_ = ctx.ShouldBindJSON(&param)
	res := c.service.GetList(param.Page, param.Size)
	base.NewResponse(ctx).Success(res)
}
