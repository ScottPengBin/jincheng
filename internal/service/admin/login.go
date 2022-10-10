package admin

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"jincheng/config"
	"jincheng/internal/core/db"
	"jincheng/internal/core/jwt"
	"jincheng/internal/model"
)

type LoginService struct {
	db *db.DataBase
	c  *config.Config
}

func NewLoginService(db *db.DataBase, c *config.Config) *LoginService {
	return &LoginService{
		db: db,
		c:  c,
	}
}

type res struct {
	*model.Users
	Token string `json:"token"`
}

func (s *LoginService) Login(args ...string) (*res, error) {

	var account *model.Users
	var r res
	s.db.Salve.Model(model.Users{}).
		Where("enabled = ?", 1).
		Where("account = ?", args[0]).
		First(&account)

	r.Users = account
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
	return &r, nil
}
