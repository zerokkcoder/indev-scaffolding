package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/pkg/response"
	"github.com/zerokkcoder/indevsca/internal/pkg/utils"
)

// Auth JWT认证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, response.ErrUnauthorized)
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Error(c, response.ErrInvalidToken)
			c.Abort()
			return
		}

		// 解析Token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.Error(c, response.ErrInvalidToken)
			c.Abort()
			return
		}

		// 将用户信息存储到上下文
		c.Set("userId", claims.UserID)
		c.Next()
	}
}
