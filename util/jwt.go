package util

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wikankun/user-service/config"
)

type jwtClaims struct {
	jwt.StandardClaims
	ID int `json:"id"`
}

func NewJWTToken(id int) (string, time.Time, error) {
	tokenExpire, _ := strconv.Atoi(config.Config.App.TokenExpire)
	expireTime := time.Now().Add(time.Duration(tokenExpire) * time.Minute)
	claims := jwtClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   "warmup-ref",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime.Unix(),
		},
		ID: id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt_secret := config.Config.App.JWTSecret
	tokenSigned, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		return "", time.Now(), err
	}
	return tokenSigned, expireTime, nil
}
