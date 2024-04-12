package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(jwtType string, secret string, userId int, role int, duration int) (string, int64, int64) {
	var exp int64
	//一种用于刷新 一种用于访问
	jwtSecret := []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)
	iat := time.Now().Unix()
	if jwtType == "refresh" {
		if duration != 0 {
			exp = time.Now().Add(time.Hour * time.Duration(duration)).Unix()
		} else {
			exp = time.Now().Add(time.Hour * 720).Unix()
		}
	} else {
		if duration != 0 {
			exp = time.Now().Add(time.Hour * time.Duration(duration)).Unix()
		} else {
			exp = time.Now().Add(time.Hour * 720).Unix()
		}
	}
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["userId"] = userId
	claims["role"] = role
	claims["type"] = jwtType
	token.Claims = claims
	tokenString, _ := token.SignedString(jwtSecret)
	fmt.Println(tokenString)
	//额外返回创建时间和到期时间
	return tokenString, exp, iat
}

func ParseToken(tokenString string, secret string) (jwt.MapClaims, error) {
	//解析JWT 如果到期 一样返回err
	jwtSecret := []byte(secret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
