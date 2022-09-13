package memberSer

import (
	"jincheng/internal/core/db"
	"jincheng/internal/model"
)

type Service struct {
	db db.DataBase
}

func NewService(db db.DataBase) Service {
	return Service{
		db: db,
	}
}

func (s *Service) GetList(page, size int) ([]model.Member, int64) {

	var item []model.Member
	var total int64

	builder := s.db.Salve.Model(&model.Member{}).
		Select("*").
		Offset((page - 1) * size).
		Limit(size)

	builder.Count(&total)

	builder.Scan(&item)

	return item, total
}
