package middlewares

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/infra/logger"
)

// Recovery 恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 检查连接是否已断开
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				// 获取堆栈信息
				stack := string(debug.Stack())
				httpRequest := fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.String())
				headers := strings.Builder{}
				for k, v := range c.Request.Header {
					headers.WriteString(fmt.Sprintf("%s: %s\n", k, v))
				}

				// 记录错误日志
				logger.Error("Recovery from panic",
					"error", err,
					"request", httpRequest,
					"headers", headers.String(),
					"stack", stack,
				)

				// 如果是断开的连接，直接返回
				if brokenPipe {
					c.Error(err.(error))
					c.Abort()
					return
				}

				// 返回 500 错误
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}
