package model

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"sync"
)

var (
	dbSingleton sync.Once
	db          *gorm.DB
	configFile  = flag.String("f", "application.yaml", "the config file")
)

/*连接参数*/
type Connector struct {
	User   string `json:"user" yaml:"user"`
	Pwd    string `json:"pwd" yaml:"pwd"`
	Host   string `json:"host" yaml:"host"`
	DbName string `json:"db" yaml:"db"`
	Port   int    `json:"port" yaml:"port"`
	Engine string `json:"engine" yaml:"engine"`
	Id     string `json:"id" yaml:"id"`
}
type DataSource struct {
	ConnectSource Connector `yaml:"datasource"`
}

func GetConnectParam() Connector {
	flag.Parse()
	yamlFile, e := ioutil.ReadFile(*configFile)
	if e != nil {
		fmt.Println(e)
	}
	var sqlDataSource DataSource
	err := yaml.Unmarshal(yamlFile, &sqlDataSource)
	if err != nil {
		fmt.Println(err)
	}
	return sqlDataSource.ConnectSource
}
func initMysqlSource(con Connector) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", con.User, con.Pwd, con.Host, con.Port, con.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}
func GetDataSource() *gorm.DB {
	dbSingleton.Do(func() {
		db = initMysqlSource(GetConnectParam())
	})
	return db
}
