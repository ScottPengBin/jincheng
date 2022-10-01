package model

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type CarInfo struct {
	gorm.Model
	Name       string                `json:"car_name" gorm:"name" binding:"required"`
	CarNo      string                `json:"car_no" binding:"required"`
	Color      string                `json:"car_color" binding:"required"`
	Note       string                `json:"car_note"`
	DeleteFlag soft_delete.DeletedAt `json:"delete_flag" gorm:"softDelete:flag,DeletedAtField:DeletedAt" gorm:"delete_flag"`
}
