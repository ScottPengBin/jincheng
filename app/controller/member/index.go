package member

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/wire"
	"jincheng/app/request/meber"
	"jincheng/internal/core/base"
	"jincheng/internal/core/valida"
	"jincheng/internal/model"
	"jincheng/internal/service/member"
	"strconv"
)

var Provider = wire.NewSet(NewController, memberSer.NewService)

func NewController(s *memberSer.MemService) *Controller {
	return &Controller{
		service: s,
	}
}

// Controller 会员中心
type Controller struct {
	service *memberSer.MemService
}

//Edit 编辑
func (c *Controller) Edit(ctx *gin.Context) {

	output := base.NewResponse(ctx)
	type reqParam struct {
		Id       uint   `json:"member_id" binding:"required" msg:"会员member_id不能为空"`
		Name     string `json:"member_name" binding:"required" msg:"会员名不能为空"`
		Mobile   string `json:"mobile" binding:"required" msg:"会员电话号码不能为空"`
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

	output := base.NewResponse(ctx)

	var memReq meber.MemRequest
	memReq.PageNum, _ = strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	memReq.PageSize, _ = strconv.Atoi(ctx.DefaultQuery("PageSize", "20"))
	memReq.Name = ctx.Query("name")
	memReq.Mobile = ctx.Query("mobile")
	memReq.CreatedAt = ctx.Query("created_at")

	res, total := c.service.GetList(&memReq)

	output.Paginate(res, total, &memReq)
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
