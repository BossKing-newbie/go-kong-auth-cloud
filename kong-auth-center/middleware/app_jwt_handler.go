package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-kong-auth-practice/kong-auth-center/config"
	"go-kong-auth-practice/kong-auth-center/constants"
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

func NewAuthJwt() *AuthJwt {
	body := config.GetJwt()
	return &AuthJwt{
		Key:         body.Key,
		ExpiredTime: body.ExpiredTime,
	}
}
func (authJwt *AuthJwt) CreateToken(userId int, loginName string) (string, error) {

	expireTime := time.Second * time.Duration(authJwt.ExpiredTime)
	claims := &CustomClaims{
		&jwt.StandardClaims{

			ExpiresAt: time.Now().Add(expireTime).Unix(),
		},
		CustomerInfo{
			UserId:    userId,
			LoginName: loginName,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(authJwt.Key))
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
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMiddleware := NewAuthJwt()
		resp := constants.NewResultBody(c)
		token := c.GetHeader("authorization")
		if len(token) == 0 {
			resp.AuthFailed("未鉴权", nil)
			c.Abort()
			return
		}
		claims, err := jwtMiddleware.ParseToken(token)
		if err != nil || time.Now().Unix() > claims.ExpiresAt {
			resp.AuthFailed("已过期", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
