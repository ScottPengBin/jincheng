package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"jincheng/config"
	"strings"
	"time"
)

type Claims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenerateToken(id uint, name string, conf *config.Config) (string, error) {
	cl := Claims{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(conf.Jwt.ExpiresAt).Unix(),
			Issuer:    conf.Jwt.Issuer,
		},
	}

	m := jwt.GetSigningMethod(strings.ToUpper(conf.Jwt.Alg))

	t := jwt.NewWithClaims(m, cl)

	return t.SigningString()
}

func ParseToken(token string, conf *config.Config) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err

}
