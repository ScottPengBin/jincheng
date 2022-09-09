package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/app/controller/pers_amount_total"
	"jincheng/app/controller/user"
)

type OptionsController struct {
	PersAmountTotal pers_amount_total.Controller
	User            user.Controller
}

var Provider = wire.NewSet(wire.Struct(new(OptionsController), "*"))

// Router 路由
func Router(oc OptionsController) func(r *gin.Engine) {
	return func(g *gin.Engine) {
		pat := g.Group("/api/zlb/riskIdentify/")
		{
			pat.GET("personalTotalAmountList", oc.PersAmountTotal.LargeAmountWarning)
		}
		jc := g.Group("/api/jc/users")
		{
			jc.GET("getList", oc.User.GetList)
		}
	}

}
