package pers_amount_total

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/internal/db"
)

var NewController = wire.NewSet(New, NewRouter)

type InitControllers func(r *gin.Engine)

func NewRouter(c Controller) InitControllers {
	return func(g *gin.Engine) {
		pat := g.Group("/api/zlb/riskIdentify/")
		{
			pat.GET("personalTotalAmountList", c.LargeAmountWarning)
		}
	}

}

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
		TotalPayAmt float64    `json:"total_pay_amt"`
		ParkId      int    `json:"park_id"`
		CardNumber  string `json:"card_number"`
		Realname    string `json:"realname"`
		Uid         int    `json:"uid"`
	}



	ctx.JSON(200, resData)
}
