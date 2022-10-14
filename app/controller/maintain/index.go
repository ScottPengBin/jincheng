package maintain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	maintain2 "jincheng/app/request/maintain"
	"jincheng/internal/core/base"
	"jincheng/internal/core/valida"
	"jincheng/internal/service/maintain"
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

	c.service.Add(&param)

	output.Success("新增成功")
}
