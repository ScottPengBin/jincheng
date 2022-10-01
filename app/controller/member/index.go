package member

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/wire"
	"jincheng/internal/core/base"
	"jincheng/internal/core/valida"
	"jincheng/internal/model"
	"jincheng/internal/service/member"
)

var Provider = wire.NewSet(NewController, memberSer.NewService)

func NewController(s *memberSer.Service) *Controller {
	return &Controller{
		service: s,
	}
}

// Controller 会员中心
type Controller struct {
	service *memberSer.Service
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
func (c Controller) Add(ctx *gin.Context) {
	//会员信息
	var memParam model.Member
	//车辆信息
	var cardParam model.CarInfo

	output := base.NewResponse(ctx)

	//会员信息验证
	if err := ctx.ShouldBindBodyWith(&memParam, binding.JSON); err != nil {
		errs := valida.Trans(err)
		output.ErrorParam(errs)
		return
	}

	//车辆信息验证
	if err := ctx.ShouldBindBodyWith(&cardParam, binding.JSON); err != nil {
		errs := valida.Trans(err)
		output.ErrorParam(errs)
		return
	}

	err := c.service.Add(&memParam, &cardParam)

	if err != nil {
		output.ErrorParam(err.Error())
		return
	}
	output.Success("添加成功")

}

func (c Controller) Test(ctx *gin.Context) {
	param := &struct {
		Msg string `json:"msg"`
	}{}

	param.Msg = ctx.Query("msg")

	base.NewResponse(ctx).Success(c.service.Test(param.Msg))

}
