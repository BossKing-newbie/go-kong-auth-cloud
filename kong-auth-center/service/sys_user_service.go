package service

import (
	"errors"
	"go-kong-auth-practice/kong-auth-center/model"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

/*线程安全的单例模式*/
var (
	sysUserModel   model.SysUserModel
	modelSingleton sync.Once
)

/*自动注入*/
func autoWiredSysUserModel() model.SysUserModel {
	modelSingleton.Do(func() {
		sysUserModel = model.NewSysUserModel()
	})
	return sysUserModel
}
func Register(user model.SysUser) error {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost) //加密处理
	user.UserPassword = string(pwd)
	user.RoleId = "user"
	err := autoWiredSysUserModel().Insert(user)
	return err
}
func Login(loginName, password string) (model.SysUser, error) {
	user := autoWiredSysUserModel().SelectByUserName(loginName)
	if user != (model.SysUser{}) {
		e := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(password)) //验证（对比）
		if e == nil {
			return user, nil
		}
	}
	return model.SysUser{}, errors.New("用户名或者密码错误")
}
func GetUser(userId int) (model.SysUser, error) {
	user := autoWiredSysUserModel().SelectByUserId(userId)
	return user, nil
}
