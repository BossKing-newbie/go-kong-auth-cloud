package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-kong-auth-practice/kong-auth-center/constants"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var authJwt JwtBody

type (
	CustomerInfo struct {
		UserId    int    `json:"userId"`
		LoginName string `json:"loginName"`
	}
	CustomClaims struct {
		*jwt.StandardClaims
		CustomerInfo
	}
	JwtBody struct {
		Key         string `json:"key" yaml:"key"`
		ExpiredTime int    `json:"expiredTime" yaml:"expiredTime"` //过期时间-秒
	}
	AuthJwt struct {
		Jwt JwtBody `yaml:"jwt"`
	}
)

func SetSignedKey(yamlPath string) {
	yamlFile, e := ioutil.ReadFile(yamlPath)
	if e != nil {
		fmt.Println(e)
	}
	var jwt AuthJwt
	err := yaml.Unmarshal(yamlFile, &jwt)
	if err != nil {
		fmt.Println(err)
	}
	authJwt = jwt.Jwt
}
func NewAuthJwt() *AuthJwt {
	return &AuthJwt{
		Jwt: authJwt,
	}
}
func (authJwt *AuthJwt) CreateToken(userId int, loginName string) (string, error) {

	expireTime := time.Second * time.Duration(authJwt.Jwt.ExpiredTime)
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
	return t.SignedString([]byte(authJwt.Jwt.Key))
}
func (authJwt *AuthJwt) ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(authJwt.Jwt.Key), nil
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
