package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Member struct {
	ID           int                   `json:"id" gorm:"id"`
	Name         string                `json:"name"`
	Phone        string                `json:"phone"`
	Status       int8                  `json:"status"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdateAt     time.Time             `json:"update_at"`
	InviteMobile string                `json:"invite_mobile"`
	DeleteFlag   soft_delete.DeletedAt `json:"delete_flag" gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}

