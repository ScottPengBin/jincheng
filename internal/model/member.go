package model

import (
	"gorm.io/plugin/soft_delete"
)

type Member struct {
	ID         uint                  `json:"id"   gorm:"primarykey"`
	Name       string                `json:"member_name" gorm:"name" binding:"required"`
	WetChatId  uint                  `json:"wet_chat_id" gorm:"wet_chat_id"`
	BrithDay   string                `json:"brith_day" gorm:"brith_day"`
	Gender     string                `json:"gender" gorm:"gender"`
	Mobile     string                `json:"mobile" gorm:"mobile" binding:"required"`
	CarId      uint                  `json:"car_id" gorm:"car_id"`
	Status     int8                  `json:"status" gorm:"status"`
	Note       string                `json:"member_note" gorm:"note"`
	CreatedAt  MyTime                `json:"created_at" gorm:"created_at"`
	UpdateAt   MyTime                `json:"update_at" gorm:"update_at"`
	DeleteFlag soft_delete.DeletedAt `json:"delete_flag" gorm:"softDelete:flag,DeletedAtField:DeletedAt" gorm:"delete_flag"`
}
