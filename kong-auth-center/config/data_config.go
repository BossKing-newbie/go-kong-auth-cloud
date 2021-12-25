package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	configFile = flag.String("f", "application.yaml", "the config file")
)

/*连接参数*/
type (
	Connector struct {
		User   string `json:"user" yaml:"user"`
		Pwd    string `json:"pwd" yaml:"pwd"`
		Host   string `json:"host" yaml:"host"`
		DbName string `json:"db" yaml:"db"`
		Port   int    `json:"port" yaml:"port"`
		Engine string `json:"engine" yaml:"engine"`
	}
	DataSource struct {
		ConnectSource Connector `yaml:"datasource"`
	}
)

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

type (
	JwtBody struct {
		Key         string `json:"key" yaml:"key"`
		ExpiredTime int    `json:"expiredTime" yaml:"expiredTime"` //过期时间-秒
	}
	AuthBody struct {
		Jwt JwtBody `yaml:"jwt"`
	}
)

func GetJwt() JwtBody {
	yamlFile, e := ioutil.ReadFile(*configFile)
	if e != nil {
		fmt.Println(e)
	}
	var jwt AuthBody
	err := yaml.Unmarshal(yamlFile, &jwt)
	if err != nil {
		fmt.Println(err)
	}
	return jwt.Jwt
}
