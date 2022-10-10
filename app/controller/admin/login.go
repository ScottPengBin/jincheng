package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/internal/core/base"
	"jincheng/internal/core/valida"
	"jincheng/internal/service/admin"
)

var LoginProvider = wire.NewSet(NewLoginController, admin.NewLoginService)

type LoginController struct {
	service *admin.LoginService
}

func NewLoginController(service *admin.LoginService) *LoginController {
	return &LoginController{
		service: service,
	}
}

// Login 登录
func (c *LoginController) Login(ctx *gin.Context) {
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
