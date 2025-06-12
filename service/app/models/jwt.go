package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/sonhineboy/gsadmin/service/global"
)

type JwtUser struct {
	Id         uint
	Name       string
	Version    int
	Roles      []string
	ApiList    map[string]string
	Permission []string
}

type CustomClaims struct {
	JwtUser
	jwt.StandardClaims
}

func (c JwtUser) NewJwtUser(id uint, name string, roles []string, apiList map[string]string, permission []string, version int) JwtUser {
	return JwtUser{
		Id:         id,
		Name:       name,
		Roles:      roles,
		ApiList:    apiList,
		Permission: permission,
		Version:    version,
	}
}

func GenToken(user JwtUser, Secret string) (string, error) {
	MySecret := []byte(Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Second * time.Duration(global.Config.MyJwt.ExpiresAt))),
			Issuer:    "ginScuiadmin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// ParseToken 解析 token
func ParseToken(tokenStr string, Secret string) (*CustomClaims, error) {
	MySecret := []byte(Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		fmt.Println(" token parse err:", err)
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	var dbVersion int
	err = global.Db.Model(&AdminUser{}).Select("version").Where("id = ?", claims.Id).Scan(&dbVersion).Error
	if err != nil || dbVersion != claims.Version {
		return nil, errors.New("invalid version")
	}

	return claims, nil
}

// RefreshToken 刷新 Token
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
		claims.StandardClaims.ExpiresAt = jwt.At(time.Now().Add(time.Second * time.Duration(global.Config.MyJwt.ExpiresAt)))
		return GenToken(claims.JwtUser, Secret)
	}
	return "", errors.New("cloud handle this token")
}
