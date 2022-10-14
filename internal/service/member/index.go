package memberSer

import (
	"errors"
	"fmt"
	"jincheng/app/request/meber"
	"jincheng/internal/core/db"
	"jincheng/internal/model"
	"time"
)

type MemService struct {
	db *db.DataBase
}

//列表
type memList struct {
	model.Member
	Age int `json:"age"`
}

//一条
type memOne struct {
	Member  model.Member  `json:"member"`
	CarInfo model.CarInfo `json:"car_info"`
}

func NewService(db *db.DataBase) *MemService {
	return &MemService{
		db: db,
	}
}

// Add 添加会员
func (s *MemService) Add(data *meber.AddReq) error {
	t := time.Now()

	tx := s.db.Master.Begin()
	//车辆信息添加
	var car model.CarInfo
	car.CreatedAt = model.MyTime(t)
	car.UpdatedAt = model.MyTime(t)
	car.CarNo = data.CarInfo.CarCommon.CarNo
	car.CarName = data.CarInfo.CarCommon.CarName
	car.CarColor = data.CarInfo.CarCommon.CarColor
	car.CarNote = data.CarInfo.CarCommon.CarNote
	carRes := tx.Model(model.CarInfo{}).Create(&car)
	if carRes.Error != nil {
		tx.Rollback()
		return carRes.Error
	}

	//会员信息添加
	var mem model.Member
	mem.CarId = car.ID
	mem.CreatedAt = model.MyTime(t)
	mem.UpdatedAt = model.MyTime(t)
	mem.Name = data.Member.MemCommon.Name
	mem.Mobile = data.Member.MemCommon.Mobile
	mem.Gender = data.Member.MemCommon.Gender
	mem.BrithDay = data.Member.MemCommon.BrithDay
	mem.MemberNote = data.Member.MemCommon.MemberNote
	memRes := tx.Model(model.Member{}).Create(&mem)
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

	//age计算
	for i, _ := range resList {
		bt, _ := time.Parse("2006-01-02", resList[i].Member.BrithDay)

		age := time.Now().Year() - bt.Year()

		//月份小于 没有到生日
		if time.Now().Month() < bt.Month() {
			age--
		}

		//月份到了 天数还没有到
		if time.Now().Month() == bt.Month() && time.Now().Day() < bt.Day() {
			age--
		}

		resList[i].Age = age

	}

	return &resList, total
}

//GetOne 获取一条
func (s *MemService) GetOne(memId int) *memOne {
	var member model.Member
	var carInfo model.CarInfo
	var memberOne memOne

	s.db.Salve.First(&member, "id = ?", memId)
	s.db.Salve.First(&carInfo, "id = ?", member.CarId)
	memberOne.Member = member
	memberOne.CarInfo = carInfo

	return &memberOne
}

//GetOneByMobile 通过电话号码获取信息
func (s *MemService) GetOneByMobile(mobile string) *model.Member {
	var member model.Member
	s.db.Salve.First(&member, "mobile = ?", mobile)
	if &member == nil {
		return nil
	}
	return &member
}

// UpdateMemberById 更新会员信息
func (s *MemService) UpdateMemberById(req *meber.UpdateReq) error {
	tx := s.db.Master.Begin()

	m := tx.Model(model.Member{}).
		Where("id = ?", req.Member.Id).
		Updates(req.Member)
	if m.Error != nil {
		tx.Rollback()
		return m.Error
	}

	c := tx.Model(model.CarInfo{}).
		Where("id = ?", req.CarInfo.Id).
		Updates(req.CarInfo)
	if c.Error != nil {
		tx.Rollback()
		return c.Error
	}

	tx.Commit()
	return nil
}

func (s *MemService) Del(id int) error {
	var mem model.Member
	var car model.CarInfo

	s.db.Master.Model(model.Member{}).
		Where("id = ?", id).
		First(&mem)

	s.db.Master.Model(model.CarInfo{}).
		Where("id = ?", id).
		First(&car)

	if &mem == nil || &car == nil {
		return errors.New("数据不存在")
	}

	t := time.Now()

	mem.UpdatedAt = model.MyTime(t)
	car.UpdatedAt = model.MyTime(t)
	car.DeleteAt = model.MyTime(t)
	mem.DeleteFlag = 1
	car.DeleteFlag = 1
	tx := s.db.Master.Begin()
	if err := tx.Model(model.Member{}).Where("id = ?", id).Updates(&mem); err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	if err := tx.Model(model.CarInfo{}).Where("id = ?", mem.CarId).Updates(&car); err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	tx.Commit()
	return nil
}
