package member

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/app/request/meber"
	"jincheng/internal/core/base"
	"jincheng/internal/core/valida"
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
	var param meber.AddReq

	output := base.NewResponse(ctx)

	//会员信息验证
	if err := ctx.ShouldBindJSON(&param); err != nil {
		errs := valida.TransMsg(err, param)
		output.ErrorParam(errs)
		return
	}

	err := c.service.Add(&param)

	if err != nil {
		output.ErrorParam(err.Error())
		return
	}
	output.Success("添加成功")

}

//GetOne 获取详情
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

//GetOneByMobile 通过电话号码获取信息
func (c *Controller) GetOneByMobile(ctx *gin.Context) {
	output := base.NewResponse(ctx)

	mobile := ctx.Query("mobile")
	if mobile == "" {
		output.ErrorParam("电话号码不能为空")
		return
	}
	res := c.service.GetOneByMobile(mobile)
	output.Success(res)
}

//UpdateMemberById 通过id更新
func (c *Controller) UpdateMemberById(ctx *gin.Context) {
	output := base.NewResponse(ctx)

	var param meber.UpdateReq
	if err := ctx.ShouldBindJSON(&param); err != nil {
		errs := valida.TransMsg(err, param)
		output.ErrorParam(errs)
		return
	}

	err := c.service.UpdateMemberById(&param)
	if err != nil {
		output.ErrorParam(err.Error())
		return
	}
	output.Success("更新成功")
}

// Del 删除
func (c *Controller) Del(ctx *gin.Context) {
	output := base.NewResponse(ctx)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		output.ErrorParam("参数错误")
		return
	}

	err = c.service.Del(id)

	if err != nil {
		output.ErrorParam(err.Error())
		return
	}
	output.Success("删除成功")
}
