package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/zerokkcoder/indevsca/internal/pkg/response"
)

// RateLimiter 请求限流中间件
func RateLimiter(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			response.Error(c, response.ErrTooManyRequests)
			c.Abort()
			return
		}
		c.Next()
	}
}
