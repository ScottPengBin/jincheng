package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/internal/core/base"
	"jincheng/internal/core/valida"
	"jincheng/internal/service/admin"
)

var Provider = wire.NewSet(NewController, admin.NewService)

type Controller struct {
	service *admin.Service
}

func NewController(service *admin.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Login(ctx *gin.Context) {
	output := base.NewResponse(ctx)

	var param struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&param); err != nil {
		msg := valida.TransMsg(err, param)
		output.ErrorParam(msg)
		return
	}

	res, err := c.service.Login(param.Account, param.Password)
	if err != nil {
		output.ErrorParam(err.Error())
		return
	}

	output.Success(res)
}

func (c *Controller) GetMenus(ctx *gin.Context) {
	output := base.NewResponse(ctx)
	var param struct {
		UserId int `json:"userId" binding:"required"`
	}


	if err := ctx.ShouldBindUri(&param); err != nil {
		msg := valida.TransMsg(err, param)
		output.ErrorParam(msg)
		return
	}
	c.service.GetMenus(param.UserId)

	output.Success("")
}