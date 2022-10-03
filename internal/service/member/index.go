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

type memList struct {
	Mem model.Member    `json:"member"`
	Car []model.CarInfo `json:"car"`
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

// GetList 列表
func (s *Service) GetList(page, size int) ([]memList, int64) {

	var memItem []model.Member
	var ResList []memList
	var total int64

	//会员信息
	builder := s.db.Salve.Model(&model.Member{}).
		Select("*")

	builder.Count(&total)

	builder.Offset((page - 1) * size).
		Limit(size).
		Scan(&memItem)

	//车辆信息
	var carIds []uint
	for _, member := range memItem {
		carIds = append(carIds, member.CarId)
	}
	var cars []model.CarInfo
	s.db.Salve.Model(model.CarInfo{}).
		Where("id in (?)", carIds).
		Select("*").
		Scan(&cars)

	//将车辆信息赋值给会员
	for i := range memItem {
		var tmp memList
		var c []model.CarInfo
		tmp.Mem = memItem[i]

		for _, car := range cars {
			if car.ID == tmp.Mem.CarId {
				c = append(c, car)
			}
		}
		tmp.Car = c
		ResList = append(ResList, tmp)
	}

	return ResList, total
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
