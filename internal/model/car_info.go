package model

import (
	"gorm.io/plugin/soft_delete"
)

type CarInfo struct {
	ID         uint                  `json:"id" gorm:"primarykey"`
	CreatedAt  MyTime                `json:"created_at"`
	UpdatedAt  MyTime                `json:"updated_at"`
	Name       string                `json:"car_name" gorm:"name" binding:"required"`
	CarNo      string                `json:"car_no" binding:"required"`
	Color      string                `json:"car_color" binding:"required"`
	Note       string                `json:"car_note"`
	DeleteFlag soft_delete.DeletedAt `json:"delete_flag" gorm:"softDelete:flag,DeletedAtField:DeletedAt" gorm:"delete_flag"`
}
