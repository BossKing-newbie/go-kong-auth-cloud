package model

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type SysUser struct {
	UserId       int       `gorm:"column:user_id" json:"userId"`
	LoginName    string    `gorm:"column:login_name" json:"loginName" binding:"required"`       //用户登录名
	UserPassword string    `gorm:"column:user_password" json:"userPassword" binding:"required"` //用户密码
	CreatedAt    time.Time `gorm:"column:insert_time" json:"insertTime"`                        //创建时间
	UpdatedAt    time.Time `gorm:"column:update_time" json:"updateTime"`                        //更新时间
	IsDeleted    string    `gorm:"column:is_deleted" json:"isDeleted"`
	RoleId       string    `gorm:"column:role_id" json:"roleId"` //数据权限ID
}

func (SysUser) TableName() string {
	return "sys_user"
}

type (
	SysUserModel interface {
		Insert(user SysUser) (err error)
		Update(user SysUser)
		SelectAll() []SysUser
		SelectByUserId(userId int) SysUser
		SelectByUserName(loginName string) SysUser
	}
	defaultSysUserModel struct {
		conn  *gorm.DB
		table string
	}
)

func (d defaultSysUserModel) Insert(user SysUser) (err error) {
	result := d.conn.Create(&user)
	return result.Error
}
func (user *SysUser) BeforeCreate(con *gorm.DB) (err error) {
	userModel := NewSysUserModel()
	userInfo := userModel.SelectByUserName(user.LoginName)
	if reflect.DeepEqual(userInfo, SysUser{}) {
		user.IsDeleted = "N"
		return
	} else {
		return errors.New("已存在该用户")
	}
}
func (d defaultSysUserModel) Update(user SysUser) {

}
func (d defaultSysUserModel) SelectAll() []SysUser {
	var sysUserList []SysUser
	d.conn.Where("is_deleted = 'N'").Find(&sysUserList)
	return sysUserList
}
func (d defaultSysUserModel) SelectByUserId(userId int) SysUser {
	var user SysUser
	d.conn.First(&user, "user_id=?", userId)
	user.UserPassword = ""
	user.UserId = -1
	return user
}
func (d defaultSysUserModel) SelectByUserName(loginName string) SysUser {
	var user SysUser
	d.conn.Where("login_name = ?", loginName).First(&user)
	return user
}
func NewSysUserModel() SysUserModel {
	return defaultSysUserModel{
		table: "sys_user",
		conn:  GetDataSource(),
	}
}
