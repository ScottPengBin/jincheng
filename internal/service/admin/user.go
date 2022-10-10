package admin

import (
	"jincheng/app/request/admin"
	"jincheng/internal/core/db"
	"jincheng/internal/model"
)

type UserService struct {
	db *db.DataBase
}

func NewUserService(db *db.DataBase) *UserService {
	return &UserService{
		db: db,
	}
}

// QueryUsersByPage 查询列表
func (s *UserService) QueryUsersByPage(param *admin.UserSearchParam) (*[]model.Users, int64) {
	tx := s.db.Salve.Model(model.Users{})

	//电话号码
	if param.Mobile != "" {
		tx.Where("mobile like ?", "%"+param.Mobile+"%")
	}

	//账号
	if param.Account != "" {
		tx.Where("account like ?", "%"+param.Account+"%")
	}

	if param.Name != "" {
		tx.Where("name like ?", "%"+param.Name+"%")
	}

	var total int64

	tx.Count(&total)

	offset := (param.PageNum - 1) * param.PageSize

	var users []model.Users
	tx.Offset(offset).Limit(param.PageSize).Scan(&users)
	return &users, total
}
