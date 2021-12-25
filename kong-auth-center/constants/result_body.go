package constants

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type (
	HttpResult interface {
		Success(msg string, data interface{})
		Failed(msg string, data interface{})
		AuthFailed(msg string, data interface{})
	}
	defaultResultBody struct {
		c *gin.Context
	}
)

func (d defaultResultBody) Success(msg string, data interface{}) {
	d.c.JSON(http.StatusOK, gin.H{
		"data": data,
		"msg":  msg,
	})
	return
}
func (d defaultResultBody) Failed(msg string, data interface{}) {
	d.c.JSON(http.StatusInternalServerError, gin.H{
		"data": data,
		"msg":  msg,
	})
	return
}
func (d defaultResultBody) AuthFailed(msg string, data interface{}) {
	d.c.JSON(http.StatusUnauthorized, gin.H{
		"data": data,
		"msg":  msg,
	})
}
func NewResultBody(c *gin.Context) HttpResult {
	return defaultResultBody{
		c: c,
	}
}
