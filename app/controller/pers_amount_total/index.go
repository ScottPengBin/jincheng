package pers_amount_total

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/internal/db"
)

var Provider = wire.NewSet(New)

func New(db db.DataBase) Controller {
	return Controller{
		db: db,
	}
}

type Controller struct {
	db db.DataBase
}

func (c Controller) LargeAmountWarning(ctx *gin.Context) {

	var resData []struct {
		TotalPayAmt float64 `json:"total_pay_amt"`
		ParkId      int     `json:"park_id"`
		CardNumber  string  `json:"card_number"`
		Realname    string  `json:"realname"`
		Uid         int     `json:"uid"`
	}

	c.db.Master.Table("zlb_task_payment_log").
		Joins("join zlb_professional on zlb_professional.uid = zlb_task_payment_log.rec_uid and zlb_professional.delete_flag = 0").
		Where("zlb_task_payment_log.pay_status = ? and zlb_task_payment_log.delete_flag = ?", 1, 0).
		Select("sum(zlb_task_payment_log.pay_amt) as total_pay_amt,zlb_task_payment_log.park_id,zlb_professional.card_number,zlb_professional.realname,zlb_professional.uid").
		Group("zlb_professional.card_number,zlb_task_payment_log.park_id").
		Offset(0).
		Limit(20).
		Scan(&resData)

	ctx.JSON(200, resData)
}
