package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/logger"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}
		c.Next()
		// 结束时间
		end := time.Now()
		// 执行时间
		latency := end.Sub(start)

		// 获取错误信息
		var errs []string
		for _, err := range c.Errors {
			errs = append(errs, err.Error())
		}

		// 记录日志
		log := logger.With(
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", path,
			"ip", c.ClientIP(),
			"latency", latency,
			// "user-agent", c.Request.UserAgent(),
		)
		if len(errs) > 0 {
			log = log.With("errors", errs)
		}

		statusCode := c.Writer.Status()
		switch {
		case statusCode >= 500:
			log.Error("Server Error")
		case statusCode >= 400:
			log.Warn("Client Error")
		case statusCode >= 300:
			log.Info("Redirection")
		default:
			log.Info("Success")
		}
	}
}
