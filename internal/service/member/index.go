package memberSer

import (
	"fmt"
	"jincheng/internal/core/db"
	"jincheng/internal/model"
	"net"
	"sync/atomic"
	"time"
)

type Service struct {
	db *db.DataBase
}

func NewService(db *db.DataBase) *Service {
	return &Service{
		db: db,
	}
}

// Add 添加会员
func (s *Service) Add(mem *model.Member, car *model.CarInfo) error {
	t := time.Now()

	tx := s.db.Master.Begin()
	car.CreatedAt = t
	car.UpdatedAt = t
	//车辆信息添加
	carRes := tx.Create(car)
	if carRes.Error != nil {
		tx.Rollback()
		return carRes.Error
	}

	//会员信息添加
	mem.CarId = car.ID
	mem.CreatedAt = t
	mem.UpdateAt = t
	memRes := tx.Create(mem)
	if memRes.Error != nil {
		tx.Rollback()
		return memRes.Error
	}
	tx.Commit()
	return nil

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

var failedNum int64
var num = int64(1)

func (s Service) Test(msg string) int {

	defer func() {
		if err := recover(); err != nil {
			atomic.AddInt64(&failedNum, 1)
			fmt.Println("failedNum=", failedNum)
		}
	}()

	if failedNum > 0 {
		atomic.AddInt64(&failedNum, -1)
		return s.Test(msg)
	}

	conn, _ := net.Dial("unix", "/tmp/pb.sock")

	l, _ := conn.Write([]byte(msg))

	defer func() {
		_ = conn.Close()
		atomic.AddInt64(&num, 1)
	}()

	fmt.Println(num)

	return l
}
