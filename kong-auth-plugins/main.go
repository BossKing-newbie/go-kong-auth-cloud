package main

import (
	"fmt"
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"log"
)

/*连接参数*/
type Connector struct {
	User   string `json:"user" yaml:"user"`
	Pwd    string `json:"pwd" yaml:"pwd"`
	Host   string `json:"host" yaml:"host"`
	DbName string `json:"db" yaml:"db"`
	Port   int    `json:"port" yaml:"port"`
	Engine string `json:"engine" yaml:"engine"`
}
type DataSourceConfig struct {
	DataSource Connector
}

func main() {
	server.StartServer(New, Version, Priority)
}

var Version = "1.0"
var Priority = 1

type Config struct {
	Message string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	host, err := kong.Request.GetHeader("host")
	if err != nil {
		log.Printf("Error reading 'host' header: %s", err.Error())
	}

	message := conf.Message
	if message == "" {
		message = "hello"
	}
	kong.Response.SetHeader("x-hello-from-go", fmt.Sprintf("Go says %s to %s", message, host))
}
