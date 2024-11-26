package context

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/domain/entity"
)

// GetUserID 从上下文获取用户ID
func GetUserID(c *gin.Context) uint {
	id, exists := c.Get("userId")
	if !exists {
		return 0
	}
	return id.(uint)
}

// GetUser 从上下文获取用户信息
func GetUser(c *gin.Context) *entity.User {
	user, exists := c.Get("user")
	if !exists {
		return nil
	}
	return user.(*entity.User)
}

// GetAdmin 从上下文获取管理员信息
func GetAdmin(c *gin.Context) *entity.Admin {
	admin, exists := c.Get("admin")
	if !exists {
		return nil
	}
	return admin.(*entity.Admin)
}

// GetTraceID 从上下文获取追踪ID
func GetTraceID(c *gin.Context) string {
	traceID, exists := c.Get("traceId")
	if !exists {
		return ""
	}
	return traceID.(string)
}
