package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Member struct {
	ID           int                   `json:"id"  gorm:"id"`
	Name         string                `json:"name" gorm:"name" binding:"required"`
	WetChat      string                `json:"wet_chat" gorm:"wet_chat" binding:"required"`
	BrithDay     string                `json:"brith_day" gorm:"brith_day"`
	Gender       string                `json:"gender" gorm:"gender"`
	Phone        string                `json:"phone" gorm:"phone" binding:"required"`
	CardNumber   string                `json:"card_number" gorm:"card_number"`
	Status       int8                  `json:"status" gorm:"status"`
	CreatedAt    time.Time             `json:"created_at" gorm:"created_at"`
	UpdateAt     time.Time             `json:"update_at" gorm:"update_at"`
	InviteMobile string                `json:"invite_mobile" gorm:"invite_mobile"`
	DeleteFlag   soft_delete.DeletedAt `json:"delete_flag" gorm:"softDelete:flag,DeletedAtField:DeletedAt" gorm:"delete_flag"`
}
