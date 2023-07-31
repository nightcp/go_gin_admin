package middlewares

import (
	"admin/core"
	"admin/core/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// ErrorHandler 错误处理
func ErrorHandler(c *gin.Context) {
	defer func() {
		resp := response.Resp{}
		r := recover()
		if r != nil {
			err := json.Unmarshal([]byte(r.(string)), &resp)
			if err != nil {
				c.JSON(response.HttpServerError, response.Resp{
					Code:   response.HttpServerError,
					Status: false,
					Msg:    core.Lang.GetMassage(c, "ServerError"),
				})
				c.Abort()
			}
			c.Header("Accept-Language", c.GetHeader("Accept-Language"))
			c.JSON(resp.Code, resp)
			c.Abort()
		}
	}()
	c.Next()
}
