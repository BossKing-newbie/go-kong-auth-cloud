package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"go-kong-auth-practice/kong-auth-plugins/constant"
	"go-kong-auth-practice/kong-auth-plugins/middleware"
)

var Version = "1.0"
var Priority = 1

type Config struct {
	DataSource Connector `yaml:"datasource"`
	Jwt        JwtBody
}
type (
	Connector struct {
		User   string `json:"user" yaml:"user"`
		Pwd    string `json:"pwd" yaml:"pwd"`
		Host   string `json:"host" yaml:"host"`
		DbName string `json:"db" yaml:"db"`
		Port   int    `json:"port" yaml:"port"`
		Engine string `json:"engine" yaml:"engine"`
	}
	JwtBody struct {
		Key         string `json:"key" yaml:"key"`
		ExpiredTime int    `json:"expiredTime" yaml:"expiredTime"` //过期时间-秒
	}
)

func main() {
	server.StartServer(New, Version, Priority)
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	/*注意事项：不要使用err.Error（）.该函数会导致插件莫名崩溃进而退出*/
	exitBody := constant.NewExitBody(kong)
	token, err := kong.Request.GetHeader("authorization")
	if err != nil {
		exitBody.Failed("token获取失败")
	}
	/*token验证*/
	authJwt := middleware.NewAuthJwt(conf.Jwt.Key, conf.Jwt.ExpiredTime)
	authErr := authJwt.Auth(token)
	if authErr != nil {
		exitBody.AuthFailed("权限验证失败")
	}
}
