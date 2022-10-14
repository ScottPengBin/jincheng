package model

type MaintainRecord struct {
	Id              int     `json:"id"`
	CarId           int     `json:"car_id"`
	MemberId        int     `json:"member_id"`
	Operator        int     `json:"operator"`
	MaintainProject string  `json:"maintain_project"`
	MaintainMoney   float64 `json:"maintain_money"`
	MaintainNote    string  `json:"maintain_note"`
	CarInfo         string  `json:"car_info"`
	MemberInfo      string  `json:"member_info"`
	MaintainBeginAt MyTime  `json:"maintain_begin_at"`
	MaintainEndAt   MyTime  `json:"maintain_end_at"`
	CreatedAt       MyTime  `json:"created_at"`
}
