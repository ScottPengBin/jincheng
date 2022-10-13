package model

import (
	"gorm.io/plugin/soft_delete"
)

type Member struct {
	ID         int                   `json:"id"   gorm:"primarykey"`
	Name       string                `json:"member_name" gorm:"name" binding:"required" msg:"会员姓名不能为空"`
	WetChatId  uint                  `json:"wet_chat_id" gorm:"wet_chat_id"`
	BrithDay   string                `json:"brith_day" gorm:"brith_day"`
	Gender     string                `json:"gender" gorm:"gender"`
	Mobile     string                `json:"mobile" gorm:"mobile" binding:"required"`
	CarId      int                   `json:"car_id" gorm:"car_id"`
	Status     int8                  `json:"status" gorm:"status"`
	MemberNote string                `json:"member_note" gorm:"note"`
	CreatedAt  MyTime                `json:"created_at" gorm:"created_at"`
	UpdatedAt  MyTime                `json:"updated_at" gorm:"updated_at"`
	DeleteFlag soft_delete.DeletedAt `json:"delete_flag" gorm:"softDelete:flag,DeletedAtField:DeletedAt" gorm:"delete_flag"`
}
