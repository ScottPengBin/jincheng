package admin

import "github.com/google/wire"

type Admin struct {
	Login *LoginController
	Menus *MenusController
	User  *UserController
}

var AdmProvider = wire.NewSet(LoginProvider, MenusProvider, UserProvider, NewAdmin)

func NewAdmin(Login *LoginController, Menus *MenusController, User *UserController) *Admin {
	return &Admin{
		Login: Login,
		Menus: Menus,
		User:  User,
	}
}
