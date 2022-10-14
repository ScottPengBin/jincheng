package model

import (
	"gorm.io/plugin/soft_delete"
)

type CarInfo struct {
	ID         int                   `json:"id" gorm:"primarykey"`
	CreatedAt  MyTime                `json:"created_at"`
	UpdatedAt  MyTime                `json:"updated_at"`
	CarName    string                `json:"car_name" gorm:"car_name" binding:"required"`
	CarNo      string                `json:"car_no" binding:"required"`
	CarColor   string                `json:"car_color" binding:"required"`
	CarNote    string                `json:"car_note"`
	DeleteAt   MyTime                `json:"-" gorm:"delete_at"`
	DeleteFlag soft_delete.DeletedAt `json:"delete_flag" gorm:"softDelete:flag,DeletedAtField:DeletedAt" gorm:"delete_flag"`
}
