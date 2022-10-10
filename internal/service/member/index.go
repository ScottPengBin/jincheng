package memberSer

import (
	"fmt"
	"jincheng/app/request/meber"
	"jincheng/internal/core/db"
	"jincheng/internal/model"
	"time"
)

type MemService struct {
	db *db.DataBase
}

type memList struct {
	model.Member
	Age int             `json:"age"`
	Car []model.CarInfo `json:"car"`
}

func NewService(db *db.DataBase) *MemService {
	return &MemService{
		db: db,
	}
}

// Add 添加会员
func (s *MemService) Add(mem *model.Member, car *model.CarInfo) error {
	t := time.Now()

	tx := s.db.Master.Begin()
	car.CreatedAt = model.MyTime(t)
	car.UpdatedAt = model.MyTime(t)
	//车辆信息添加
	carRes := tx.Create(car)
	if carRes.Error != nil {
		tx.Rollback()
		return carRes.Error
	}

	//会员信息添加
	mem.CarId = car.ID
	mem.CreatedAt = model.MyTime(t)
	mem.UpdateAt = model.MyTime(t)
	memRes := tx.Create(mem)
	if memRes.Error != nil {
		tx.Rollback()
		return memRes.Error
	}
	tx.Commit()
	return nil

}

// GetList 列表
func (s *MemService) GetList(req *meber.MemRequest) (*[]memList, int64) {

	var resList []memList
	var total int64

	//会员信息
	builder := s.db.Salve.Model(&model.Member{}).
		Select("*")

	if req.Name != "" {
		builder.Where("name like ?", fmt.Sprintf("%%%s%%", req.Name))
	}

	if req.Mobile != "" {
		builder.Where("mobile like ?", fmt.Sprintf("%%%s%%", req.Mobile))
	}

	builder.Count(&total)

	builder.Offset((req.PageNum - 1) * req.PageSize).
		Limit(req.PageSize).
		Scan(&resList)

	//车辆信息
	var carIds []uint
	for _, member := range resList {
		carIds = append(carIds, member.CarId)
	}
	var cars []model.CarInfo

	if len(carIds) > 0 {
		s.db.Salve.Model(model.CarInfo{}).
			Where("id in (?)", carIds).
			Select("*").
			Scan(&cars)
	}

	//将车辆信息赋值给会员
	for i, item := range resList {

		var c []model.CarInfo
		for _, car := range cars {
			if car.ID == item.CarId {
				c = append(c, car)
			}
		}

		bt, _ := time.Parse("2006-01-02", resList[i].BrithDay)

		resList[i].Age = time.Now().AddDate(-bt.Year(), int(-bt.Month()), -bt.Day()).Year()
		resList[i].Car = c
	}

	return &resList, total
}
