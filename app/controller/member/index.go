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

//Edit 编辑
func (c *Controller) Edit(ctx *gin.Context) {

	output := base.NewResponse(ctx)
	type reqParam struct {
		Id       uint   `json:"member_id" binding:"required" msg:"会员member_id不能为空"`
		Name     string `json:"member_name" binding:"required" msg:"会员名不能为空"`
		Phone    string `json:"phone" binding:"required" msg:"会员电话号码不能为空"`
		Note     string `json:"note"`
		BrithDay string `json:"brith_day"`
		Gender   string `json:"gender"`
	}
	var param reqParam

	if err := ctx.ShouldBindJSON(&param); err != nil {
		output.ErrorParam(valida.TransMsg(err, param))
		return
	}

	output.Success(param)
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
func (c *Controller) Add(ctx *gin.Context) {
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

func (c *Controller) Test(ctx *gin.Context) {
	param := &struct {
		Msg string `json:"msg"`
	}{}

	param.Msg = ctx.Query("msg")

	base.NewResponse(ctx).Success(c.service.Test(param.Msg))

}
