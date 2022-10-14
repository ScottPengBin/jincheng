package maintain

import "jincheng/internal/model"

type ListReq struct {
	PageNum  int
	PageSize int
}

func (r *ListReq) GetPageNum() int {
	return r.PageNum
}

func (r *ListReq) GetPageSize() int {
	return r.PageSize
}

type AddReq struct {
	Mobile          string       `json:"mobile" binding:"required" msg:"会员电话号码不能为空"`
	MaintainProject string       `json:"maintain_project" binding:"required" msg:"保养项目不能为空"`
	MaintainMoney   float64      `json:"maintain_money" binding:"required" msg:"保养金额不能为空"`
	MaintainBeginAt model.MyTime `json:"maintain_begin_at" binding:"required" msg:"保养开始不能为空"`
	MaintainEndAt   model.MyTime `json:"maintain_end_at" binding:"required" msg:"保养结束不能为空"`
	MaintainNote    string       `json:"maintain_note"`
}
