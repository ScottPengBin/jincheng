package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/app/controller/member"
)

type OptionsController struct {
	Member *member.Controller
}

var Provider = wire.NewSet(
	member.Provider,
	wire.Struct(new(OptionsController), "*"),
)

// Router 路由
func Router(oc *OptionsController) func(r *gin.Engine) {
	return func(g *gin.Engine) {
		jc := g.Group("/api/jc/member")
		{
			jc.GET("getList", oc.Member.GetList)
			jc.GET("test", oc.Member.Test)
			jc.POST("add", oc.Member.Add)
			jc.POST("edit", oc.Member.Edit)
		}
	}

}
