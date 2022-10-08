package admin

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"jincheng/config"
	"jincheng/internal/core/db"
	"jincheng/internal/core/jwt"
	"jincheng/internal/model"
)

type Service struct {
	db *db.DataBase
	c  *config.Config
}

func NewService(db *db.DataBase, c *config.Config) *Service {
	return &Service{
		db: db,
		c:  c,
	}
}

type res struct {
	*model.Users
	Token string `json:"token"`
}

func (s *Service) Login(args ...string) (*res, error) {

	var account *model.Users
	var r res
	s.db.Salve.Model(model.Users{}).
		Where("enabled = ?", 1).
		Where("account = ?", args[0]).
		First(&account)

	r.Users = account
	hash := sha1.New()
	hash.Write([]byte(args[1]))
	p := hex.EncodeToString(hash.Sum([]byte("")))
	if account.Password != p {
		return nil, errors.New("密码有误")
	}

	t, err := jwt.GenerateToken(account.Id, account.Name, s.c)
	if err != nil {
		return nil, err
	}
	r.Token = t
	return &r, nil
}

func (s *Service) GetMenus(uId int) *[]model.Menus {
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
