package meber

type MemRequest struct {
	PageNum   int
	PageSize  int
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	CreatedAt string `json:"created_at"`
}

func (m *MemRequest) GetPageNum() int {
	return m.PageNum
}
func (m *MemRequest) GetPageSize() int {
	return m.PageSize
}

//UpdateReq 会员更新信息
type UpdateReq struct {
	Member struct {
		Id int `json:"id" binding:"required"`
		MemCommon
	} `json:"member"`
	CarInfo struct {
		Id int `json:"id" binding:"required"`
		CarCommon
	} `json:"car_info"`
}

type MemCommon struct {
	Name       string `json:"member_name" binding:"required" msg:"会员姓名不能为空"`
	Mobile     string `json:"mobile" binding:"required" msg:"会员电话号码不能为空"`
	Gender     string `json:"gender" binding:"required" msg:"会员性别不能为空"`
	BrithDay   string `json:"brith_day" binding:"required" msg:"会员生日不能为空"`
	MemberNote string `json:"member_note"`
}

type CarCommon struct {
	CarNo    string `json:"car_no" binding:"required" msg:"车牌号不能为空"`
	CarName  string `json:"car_name" binding:"required" msg:"车辆名称不能为空"`
	CarColor string `json:"car_color" binding:"required" msg:"车辆颜色不能为空"`
	CarNote  string `json:"car_note"`
}

//AddReq 新增请求
type AddReq struct {
	Member struct {
		MemCommon
	} `json:"member"`
	CarInfo struct {
		CarCommon
	} `json:"car_info"`
}
