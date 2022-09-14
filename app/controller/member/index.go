package member

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/internal/core/base"
	"jincheng/internal/core/valida"
	"jincheng/internal/model"
	"jincheng/internal/service/member"
)

var Provider = wire.NewSet(NewController, memberSer.NewService)

func NewController(s memberSer.Service) Controller {
	return Controller{
		service: s,
	}
}

// Controller 会员中心
type Controller struct {
	service memberSer.Service
}

// GetList 获取会员列表
func (c *Controller) GetList(ctx *gin.Context) {
	var param base.ReqPaginateParam

	_ = ctx.ShouldBindJSON(&param)

	if param.Current <= 0 {
		param.Current = 1
	}

	if param.Size <= 0 {
		param.Size = 20
	}

	res, total := c.service.GetList(param.Current, param.Size)

	base.NewResponse(ctx).Paginate(res, total, param)
}

// Add 新增会员
func (c Controller) Add(ctx *gin.Context)  {
	var memParam model.Member

	if err := ctx.ShouldBindJSON(&memParam);err != nil{
		errs := valida.Trans(err)
		base.NewResponse(ctx).ErrorParam(errs)
		return
	}

}