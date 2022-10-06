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

	res := c.service.Login(param.Account, param.Password)

	output.Success(res)
}
