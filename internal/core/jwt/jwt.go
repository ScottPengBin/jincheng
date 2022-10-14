package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"jincheng/config"
	"strings"
	"time"
)

type Claims struct {
	ID        uint   `json:"id"`
	AdminName string `json:"admin_name"`
	jwt.StandardClaims
}

func GenerateToken(id uint, name string, conf *config.Config) (string, error) {
	d, err := time.ParseDuration(conf.Jwt.ExpiresAt)

	if err != nil {
		return "", errors.New("过期时间设置有误")
	}

	cl := Claims{
		ID:        id,
		AdminName: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(d).Unix(),
			Issuer:    conf.Jwt.Issuer,
		},
	}

	var signingMethodMap = map[string]jwt.SigningMethod{
		"ES256": jwt.SigningMethodES256,
		"ES384": jwt.SigningMethodES384,
		"ES512": jwt.SigningMethodES512,
		"HS256": jwt.SigningMethodHS512,
		"HS384": jwt.SigningMethodHS512,
		"HS512": jwt.SigningMethodHS512,
		"RS256": jwt.SigningMethodRS512,
		"RS384": jwt.SigningMethodRS512,
		"RS512": jwt.SigningMethodRS512,
	}

	if SigningMethod, ok := signingMethodMap[strings.ToUpper(conf.Jwt.Alg)]; ok {

		t := jwt.NewWithClaims(SigningMethod, cl)
		mySigningKey := []byte(conf.Jwt.Secret)

		return t.SignedString(mySigningKey)

	}

	return "", errors.New(conf.Jwt.Alg + "加密方式不存在")

}

func ParseToken(token string, conf *config.Config) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Jwt.Secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err

}
