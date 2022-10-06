package admin

import (
	"jincheng/internal/core/db"
	"jincheng/internal/model"
)

type Service struct {
	db *db.DataBase
}

func NewService(db *db.DataBase) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Login(args ...string) *model.Admin {
	var account *model.Admin
	s.db.Salve.Model(model.Admin{}).
		Where("account = ?", args[0]).
		First(&account)

	return account

}
