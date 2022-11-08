package models

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type JwtUser struct {
	Id    uint
	Name  string
	Roles []string
}

type CustomClaims struct {
	JwtUser
	jwt.StandardClaims
}

func GenToken(user JwtUser, Secret string) (string, error) {
	MySecret := []byte(Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 3)), //3小时过期
			Issuer:    "ginScuiadmin",                        //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// 解析 token
func ParseToken(tokenStr string, Secret string) (*CustomClaims, error) {
	MySecret := []byte(Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		fmt.Println(" token parse err:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 刷新 Token
func RefreshToken(tokenStr string, Secret string) (string, error) {
	MySecret := []byte(Secret)

	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = jwt.At(time.Now().Add(time.Minute * 10))
		return GenToken(claims.JwtUser, Secret)
	}
	return "", errors.New("cloud handle this token")
}
