package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type CarInfo struct {
	ID         uint                  `json:"id" gorm:"primarykey"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
	Name       string                `json:"car_name" gorm:"name" binding:"required"`
	CarNo      string                `json:"car_no" binding:"required"`
	Color      string                `json:"car_color" binding:"required"`
	Note       string                `json:"car_note"`
	DeleteFlag soft_delete.DeletedAt `json:"delete_flag" gorm:"softDelete:flag,DeletedAtField:DeletedAt" gorm:"delete_flag"`
}
