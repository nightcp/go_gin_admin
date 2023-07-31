package middlewares

import (
	"admin/core"
	"github.com/gin-gonic/gin"
)

// LocaleHandler 处理本地语言
func LocaleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("Accept-Language")
		if locale == "" {
			locale = core.Config.Locale
		}
		c.Request.Header.Set("Accept-Language", locale)
		c.Next()
	}
}
