package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2

// const TokenExpireDuration = time.Second * 60

var Secret = []byte("人生路漫漫")

type MyClaims struct {
	UserName string
	jwt.StandardClaims
}

// get token
func GetToken(username string) (string, error) {
	cla := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "lx-jwt",                                   // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	fmt.Println("Token = ", token)
	return token.SignedString(Secret) // 进行签名生成对应的token
}

// parse token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func main() {

	//secret :="1212121213131"
	//token,_:=models.GenToken(models.AdminUser{
	//	Name: "nihao",
	//	RealName:"afasf",
	//	Password:"asdfasdf",
	//},secret)
	//
	//fmt.Println(token)
	//ss,err:=models.ParseToken(token,secret)
	//
	//if err !=nil {
	//	fmt.Println(err.Error())
	//}

	//fmt.Println(ss)

}
