package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sunyd/go-demo1/model"
	"time"
)

var jwtKey = []byte("a_secret_ecret")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

/**
颁发token.
*/
func ReleaseToken(user *model.User) (string, error) {
	//超时时间
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	// 定义一个声明
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "learn.tech",
			Subject:   "user token",
		},
	}
	//通过声明颁发一个token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//加密一个token返回给客户端
	if tokenResult, err := token.SignedString(jwtKey); err != nil {
		return "", err
	} else {
		return tokenResult, nil
	}
}

/**
解析token.
*/
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
