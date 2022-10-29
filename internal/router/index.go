package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"jincheng/app/controller/admin"
	"jincheng/app/controller/integral"
	"jincheng/app/controller/maintain"
	"jincheng/app/controller/member"
)

type OptionsController struct {
	Member   *member.Controller
	Admin    *admin.Admin
	Maintain *maintain.Controller
	Integral *integral.Controller
}

var Provider = wire.NewSet(
	member.Provider,
	admin.AdmProvider,
	maintain.Provider,
	integral.Provider,
	wire.Struct(new(OptionsController), "*"),
)

// Router 路由
func Router(oc *OptionsController) func(r *gin.Engine) {

	return func(g *gin.Engine) {
		adm := g.Group("api/jc/admin")
		{
			//会员列表
			mem := adm.Group("member")
			{
				mem.GET("getList", oc.Member.GetList)
				mem.POST("add", oc.Member.Add)
				mem.POST("edit", oc.Member.Edit)
				mem.GET("getOne", oc.Member.GetOne)
				mem.GET("getOneByMobile", oc.Member.GetOneByMobile)
				mem.POST("updateMemberById", oc.Member.UpdateMemberById)
				mem.DELETE("delete/:id", oc.Member.Del)
			}

			//目录
			menu := adm.Group("menu")
			{
				menu.GET("queryMenus", oc.Admin.Menus.QueryMenus)
				menu.POST("addMenu", oc.Admin.Menus.AddMenus)
			}

			adm.POST("login", oc.Admin.Login.Login)
			adm.GET("authority/queryUserMenus", oc.Admin.Menus.GetMenus)
			adm.GET("user/queryUsersByPage", oc.Admin.User.QueryUsersByPage)

			//保养
			mt := adm.Group("maintain")
			{
				mt.GET("getList", oc.Maintain.GetList)
				mt.GET("getOne", oc.Maintain.GetOne)
				mt.POST("add", oc.Maintain.Add)
				mt.POST("getList", oc.Maintain.Edit)
				mt.DELETE("delete/:id", oc.Maintain.Del)
			}
		}
	}

}
