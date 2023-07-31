package middlewares

import (
	"admin/core"
	"admin/core/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

var limiter = ratelimit.NewBucketWithQuantum(time.Second, core.Config.RateLimiterCapacity, core.Config.RateLimiterQuantum)

// RateLimiterHandler 处理限流
func RateLimiterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.TakeAvailable(1) == 0 {
			response.Fail(c, "TooManyRequest", nil, response.Options{HttpCode: response.HttpForbidden})
		}
		c.Writer.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limiter.Available()))
		c.Writer.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.Capacity()))
		c.Next()
	}
}
