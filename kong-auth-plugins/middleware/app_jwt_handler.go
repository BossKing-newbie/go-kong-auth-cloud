package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type (
	CustomerInfo struct {
		UserId    int    `json:"userId"`
		LoginName string `json:"loginName"`
	}
	CustomClaims struct {
		*jwt.StandardClaims
		CustomerInfo
	}
	AuthJwt struct {
		Key         string `json:"key" yaml:"key"`
		ExpiredTime int    `json:"expiredTime" yaml:"expiredTime"` //过期时间-秒
	}
)

func NewAuthJwt(key string, expiredTime int) *AuthJwt {
	return &AuthJwt{
		Key:         key,
		ExpiredTime: expiredTime,
	}
}
func (authJwt *AuthJwt) ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(authJwt.Key), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
func (authJwt *AuthJwt) Auth(token string) error {
	if len(token) == 0 {
		err := errors.New("未鉴权")
		return err
	} else {
		claims, err := authJwt.ParseToken(token)
		if err != nil || time.Now().Unix() > claims.ExpiresAt {
			err = errors.New("已过期")
			return err
		} else {
			return nil
		}
	}
}
