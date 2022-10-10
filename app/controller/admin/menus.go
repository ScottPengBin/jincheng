package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/internal/core/base"
	"jincheng/internal/core/valida"
	"jincheng/internal/model"
	"jincheng/internal/service/admin"
	"strconv"
)

type MenusController struct {
	service *admin.MenusService
}

var MenusProvider = wire.NewSet(NewMensController, admin.NewMenusService)

func NewMensController(service *admin.MenusService) *MenusController {
	return &MenusController{
		service: service,
	}
}

// GetMenus 获取菜单
func (c *MenusController) GetMenus(ctx *gin.Context) {
	output := base.NewResponse(ctx)
	userId := ctx.Query("userId")
	uid, err := strconv.Atoi(userId)
	if err != nil {
		output.ErrorParam("userId必须为数字")
		return
	}

	res := c.service.GetMenus(uid)

	output.Success(res)
}

// QueryMenus 菜单管理
func (c *MenusController) QueryMenus(ctx *gin.Context) {
	output := base.NewResponse(ctx)

	res := c.service.QueryMenus()

	output.Success(res)
}

// AddMenus 添加菜单
func (c *MenusController) AddMenus(ctx *gin.Context) {
	output := base.NewResponse(ctx)

	var menu model.Menus
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		msg := valida.TransMsg(err, menu)
		output.ErrorParam(msg)
		return
	}

	err := c.service.AddMenus(&menu)
	if err != nil {
		output.ErrorParam(err.Error())
		return
	}

	output.Success("菜单添加成功")
}
