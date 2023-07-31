package response

import (
	"admin/core"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// Success 返回成功
func Success(c *gin.Context, msg string, data interface{}, options ...Options) {
	resp := Resp{
		Status: true,
		Msg:    msg,
		Data:   data,
	}
	respJson(c, resp, options...)
}

// Fail 返回错误
func Fail(c *gin.Context, msg string, data interface{}, options ...Options) {
	respJson(c, Resp{
		Status: false,
		Msg:    msg,
		Data:   data,
	}, options...)
}

// respJson 返回JSON
func respJson(c *gin.Context, resp Resp, options ...Options) {
	code := HttpStatusOk
	var msgArgs []interface{}
	if len(options) > 0 {
		for _, option := range options {
			if option.HttpCode > 0 {
				code = option.HttpCode
			}
			if len(option.MsgArgs) > 0 {
				msgArgs = option.MsgArgs
			}
		}
	}
	resp.Code = code
	resp.Msg = core.Lang.GetMassage(c, resp.Msg, msgArgs)
	if resp.Status == false {
		_json, _ := json.Marshal(resp)
		panic(string(_json))
	}
	c.Header("Accept-Language", c.GetHeader("Accept-Language"))
	c.JSON(code, resp)
}
