package admin

import (
	"jincheng/internal/core/db"
	"jincheng/internal/model"
	"time"
)

type MenusService struct {
	db *db.DataBase
}

func NewMenusService(db *db.DataBase) *MenusService {
	return &MenusService{
		db: db,
	}
}

// GetMenus 获取菜单
func (s *MenusService) GetMenus(uId int) *[]model.Menus {
	type role struct {
		Type   int `json:"type"`
		RoleId int `json:"role_id"`
	}

	var roleRes []role

	s.db.Salve.Table("jc_user_roles").
		Joins("left join jc_roles on jc_roles.id = jc_user_roles.roleId").
		Where("jc_user_roles.userId = ?", uId).
		Select("jc_roles.type,jc_roles.id as role_id").
		Scan(&roleRes)

	super := false
	var roleId []int
	for _, re := range roleRes {
		if re.Type == 1 {
			super = true
			break
		}
		roleId = append(roleId, re.RoleId)
	}

	var menus []model.Menus
	//超级管理员
	if super {
		s.db.Salve.Model(model.Menus{}).
			Scan(&menus)
		return &menus
	}

	//普通
	s.db.Salve.Model(model.Menus{}).
		Joins("left join jc_role_menus on jc_role_menus.menuId = jc_menus.id").
		Where("jc_role_menus.id in (?)", roleId).
		Select("jc_menus.*").
		Scan(&menus)

	return &menus
}

//QueryMenus 菜单管理
func (s *MenusService) QueryMenus() *[]model.Menus {
	var menus []model.Menus
	s.db.Salve.Model(model.Menus{}).
		Scan(&menus)
	return &menus
}

//AddMenus 添加菜单
func (s *MenusService) AddMenus(menu *model.Menus) error {
	t := time.Now()
	menu.CreatedAt = model.MyTime(t)
	menu.UpdatedAt = model.MyTime(t)
	tx := s.db.Master.Model(model.Menus{}).Create(menu)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
