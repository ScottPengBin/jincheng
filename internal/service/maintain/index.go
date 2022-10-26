package maintain

import (
	"encoding/json"
	"errors"
	"jincheng/app/request/maintain"
	"jincheng/internal/core/db"
	"jincheng/internal/model"
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

func (s *Service) GetList(param *maintain.ListReq) (*[]model.MaintainRecord, int64) {
	var list []model.MaintainRecord
	builder := s.db.Salve.Model(model.MaintainRecord{})

	var total int64
	builder.Count(&total)

	builder.Select("*").
		Offset((param.PageNum - 1) * param.PageSize).
		Limit(param.PageSize).
		Scan(&list)

	return &list, total
}

func (s *Service) Add(req *maintain.AddReq) error {
	salve := s.db.Salve

	var mem model.Member
	salve.Model(model.Member{}).
		Where("mobile = ?", req.Mobile).
		First(&mem)

	var car model.CarInfo
	salve.Model(model.CarInfo{}).
		Where("id = ?", mem.CarId).
		First(&car)

	t := time.Now()
	var record model.MaintainRecord
	jc, _ := json.Marshal(car)
	jm, _ := json.Marshal(mem)

	record.CarId = mem.CarId
	record.MemberId = mem.ID
	record.CarInfo = string(jc)
	record.MemberInfo = string(jm)
	record.CreatedAt = model.MyTime(t)
	record.MaintainBeginAt = req.MaintainBeginAt
	record.MaintainEndAt = req.MaintainEndAt

	s.db.Master.Model(model.MaintainRecord{}).Create(&record)

	return nil
}

func (s *Service) Edit(req *maintain.EditReq) error {

	tx := s.db.Master.Model(model.MaintainRecord{}).
		Where("id = ?", req.Id).
		Updates(req.AddReq)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (s *Service) GetOne(id int) *model.MaintainRecord {
	var record model.MaintainRecord
	s.db.Salve.Model(model.MaintainRecord{}).
		Where("id = ?", id).
		First(&record)
	return &record
}

func (s *Service) Del(id int) error {
	var record model.MaintainRecord

	s.db.Master.Model(model.MaintainRecord{}).
		Where("id = ?", id).
		First(&record)
	if &record == nil {
		return errors.New("数据不存在")
	}

	s.db.Master.Model(model.MaintainRecord{}).
		Where("id = ?", id).
		Delete(&record)

	return nil
}
