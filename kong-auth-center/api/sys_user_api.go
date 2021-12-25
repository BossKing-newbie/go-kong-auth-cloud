package api

import (
	"github.com/gin-gonic/gin"
	"go-kong-auth-practice/kong-auth-center/constants"
	"go-kong-auth-practice/kong-auth-center/middleware"
	"go-kong-auth-practice/kong-auth-center/model"
	"go-kong-auth-practice/kong-auth-center/service"
)

/*
认证中心只允许具备登录与注册功能
*/
func RegisterApi(c *gin.Context) {
	var user model.SysUser
	resp := constants.NewResultBody(c)
	if err := c.ShouldBindJSON(&user); err != nil {
		resp.Failed(err.Error(), nil)
		return
	} else {
		err = service.Register(user)
		if err == nil {
			resp.Success("注册成功", nil)
			return
		} else {
			resp.Failed(err.Error(), nil)
			return
		}
	}
}
func LoginApi(c *gin.Context) {
	resp := constants.NewResultBody(c)
	loginName := c.PostForm("userName")
	password := c.PostForm("password")
	jwtMiddleware := middleware.NewAuthJwt()
	if len(loginName) == 0 || len(password) == 0 {
		resp.Failed("用户名或者密码不能为空", nil)
	} else {
		user, err := service.Login(loginName, password)
		if err == nil {
			token, _ := jwtMiddleware.CreateToken(user.UserId, user.LoginName)
			resp.Success("登录成功", token)
		} else {
			resp.Failed(err.Error(), nil)
		}
	}
}
