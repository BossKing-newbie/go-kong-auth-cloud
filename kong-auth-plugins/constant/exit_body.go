package constant

import (
	"encoding/json"
	"github.com/Kong/go-pdk"
	"net/http"
)

type ResultBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type (
	HttpResult interface {
		Failed(msg string)
		AuthFailed(msg string)
	}
	defaultExitBody struct {
		kong *pdk.PDK
	}
)

func (d defaultExitBody) Failed(msg string) {
	exitMap := map[string]interface{}{
		"code": http.StatusInternalServerError,
		"msg":  msg,
	}
	exitJson, _ := json.Marshal(exitMap)
	d.kong.Response.Exit(http.StatusInternalServerError, string(exitJson), nil)
}
func (d defaultExitBody) AuthFailed(msg string) {
	exitMap := map[string]interface{}{
		"code": http.StatusUnauthorized,
		"msg":  msg,
	}
	exitJson, _ := json.Marshal(exitMap)
	d.kong.Response.Exit(http.StatusUnauthorized, string(exitJson), nil)
}
func NewExitBody(kong *pdk.PDK) HttpResult {
	kong.Response.SetHeader("Content-Type", "application/json; charset=utf-8")
	return defaultExitBody{
		kong: kong,
	}
}
