package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Member struct {
	ID         uint                  `json:"id"   gorm:"primarykey"`
	Name       string                `json:"member_name" gorm:"name" binding:"required"`
	WetChatId  uint                  `json:"wet_chat_id" gorm:"wet_chat_id"`
	BrithDay   string                `json:"brith_day" gorm:"brith_day"`
	Gender     string                `json:"gender" gorm:"gender"`
	Phone      string                `json:"phone" gorm:"phone" binding:"required"`
	CarId      uint                  `json:"car_id" gorm:"car_id"`
	Status     int8                  `json:"status" gorm:"status"`
	Note       string                `json:"member_note" gorm:"note"`
	CreatedAt  time.Time             `json:"created_at" gorm:"created_at"`
	UpdateAt   time.Time             `json:"update_at" gorm:"update_at"`
	DeleteFlag soft_delete.DeletedAt `json:"delete_flag" gorm:"softDelete:flag,DeletedAtField:DeletedAt" gorm:"delete_flag"`
}
