package user

import "jincheng/internal/db"

type Service struct {
	db db.DataBase
}

func NewService(db db.DataBase) Service {
	return Service{
		db: db,
	}
}

type Item struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (s *Service) GetList(page, size int) []Item {

	var item []Item

	s.db.Master.Table("jc_users").
		Select("*").
		Offset((page - 1) * size).
		Limit(size).
		Scan(&item)

	return item
}
