package integral

import (
	"github.com/google/wire"
	"jincheng/internal/core/db"
)

var Provider = wire.NewSet(NewService)

type Service struct {
	db *db.DataBase
}

func NewService(db *db.DataBase) *Service {
	return &Service{
		db: db,
	}
}