package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/app/controller/admin"
	"jincheng/app/controller/member"
)

type OptionsController struct {
	Member *member.Controller
	Admin  *admin.Admin
}

var Provider = wire.NewSet(
	member.Provider,
	admin.AdmProvider,
	wire.Struct(new(OptionsController), "*"),
)

// Router 路由
func Router(oc *OptionsController) func(r *gin.Engine) {

	return func(g *gin.Engine) {
		adm := g.Group("api/jc/admin")
		{
			mem := adm.Group("member")
			{
				mem.GET("getList", oc.Member.GetList)
				mem.POST("add", oc.Member.Add)
				mem.POST("edit", oc.Member.Edit)
			}

			menu := adm.Group("menu")
			{
				menu.GET("queryMenus", oc.Admin.Menus.QueryMenus)
				menu.POST("addMenu", oc.Admin.Menus.AddMenus)
			}

			adm.POST("login", oc.Admin.Login.Login)
			adm.GET("authority/queryUserMenus", oc.Admin.Menus.GetMenus)

			adm.GET("user/queryUsersByPage", oc.Admin.User.QueryUsersByPage)
		}
	}

}
