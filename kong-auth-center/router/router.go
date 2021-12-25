package router

import (
	"github.com/gin-gonic/gin"
	"go-kong-auth-practice/kong-auth-center/api"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", api.LoginApi)
	router.POST("/register", api.RegisterApi)
	return router
}
