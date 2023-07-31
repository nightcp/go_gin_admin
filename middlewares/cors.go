package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// CorsHandler 处理跨域请求
func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"OPTIONS", "GET", "POST", "POST", "DELETE", "PUT"},
			AllowHeaders: []string{"*"},
			MaxAge:       12 * time.Hour,
		})
		c.Next()
	}
}
