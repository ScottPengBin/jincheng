package admin

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"jincheng/config"
	"jincheng/internal/core/db"
	"jincheng/internal/core/jwt"
	"jincheng/internal/model"
)

type Service struct {
	db *db.DataBase
	c  *config.Config
}

func NewService(db *db.DataBase, c *config.Config) *Service {
	return &Service{
		db: db,
		c:  c,
	}
}

type res struct {
	*model.Admin
	Token string `json:"token"`
}

func (s *Service) Login(args ...string) (*res, error) {

	var account *model.Admin
	var r res
	s.db.Salve.Model(model.Admin{}).
		Where("account = ?", args[0]).
		First(&account)

	r.Admin = account
	hash := sha1.New()
	hash.Write([]byte(args[1]))
	p := hex.EncodeToString(hash.Sum([]byte("")))
	if account.Password != p {
		return nil, errors.New("密码有误")
	}

	t, err := jwt.GenerateToken(account.Id, account.Name, s.c)
	if err != nil {
		return nil, err
	}
	r.Token = t

	fmt.Println(r)
	return &r, nil

}
