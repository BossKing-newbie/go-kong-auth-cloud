package model

import (
	"fmt"
	config "go-kong-auth-practice/kong-auth-center/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	dbSingleton sync.Once
	db          *gorm.DB
)

func initMysqlSource(con config.Connector) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", con.User, con.Pwd, con.Host, con.Port, con.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
func GetDataSource() *gorm.DB {
	dbSingleton.Do(func() {
		db = initMysqlSource(config.GetConnectParam())
	})
	return db
}
