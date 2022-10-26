package maintain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	maintain2 "jincheng/app/request/maintain"
	"jincheng/internal/core/base"
	"jincheng/internal/core/valida"
	"jincheng/internal/service/maintain"
	"strconv"
)

type Controller struct {
	service *maintain.Service
}

var Provider = wire.NewSet(NewController, maintain.NewService)

func NewController(service *maintain.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// GetList 保养记录列表
func (c *Controller) GetList(ctx *gin.Context) {
	output := base.NewResponse(ctx)

	var param maintain2.ListReq

	_ = ctx.ShouldBindJSON(&param)

	res, total := c.service.GetList(&param)

	output.Paginate(res, total, &param)
}

func (c *Controller) Add(ctx *gin.Context) {
	output := base.NewResponse(ctx)
	var param maintain2.AddReq

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		msg := valida.TransMsg(err, param)
		output.ErrorParam(msg)
		return
	}

	err = c.service.Add(&param)
	if err != nil {
		output.ErrorParam(err.Error())
		return
	}

	output.Success("新增成功")
}

func (c *Controller) Edit(ctx *gin.Context) {
	output := base.NewResponse(ctx)
	//数据校验
	var editParam maintain2.EditReq
	err := ctx.ShouldBindJSON(&editParam)
	if err != nil {
		msg := valida.TransMsg(err, editParam)
		output.ErrorParam(msg)
		return
	}
	err = c.service.Edit(&editParam)
	if err != nil {
		output.ErrorParam(err.Error())
		return
	}

	output.Success("修改成功")
}

func (c *Controller) GetOne(ctx *gin.Context) {
	output := base.NewResponse(ctx)
	id, _ := strconv.Atoi(ctx.Query("id"))

	if id > 0 {
		res := c.service.GetOne(id)
		output.Success(res)
		return
	}

	output.ErrorParam("id有误")

}

func (c *Controller) Del(ctx *gin.Context)  {
	output := base.NewResponse(ctx)
	id, _ := strconv.Atoi(ctx.Query("id"))

	if id > 0 {
		res := c.service.Del(id)
		output.Success(res)
		return
	}

	output.ErrorParam("id有误")
}
