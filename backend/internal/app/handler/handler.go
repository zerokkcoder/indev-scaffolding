package handler

import "github.com/gin-gonic/gin"

// AuthHandler 认证处理器接口
type AuthHandler interface {
	Login() gin.HandlerFunc
	Logout(c *gin.Context)
}

// UserHandler 用户处理器接口
type UserHandler interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Enable(c *gin.Context)
	Disable(c *gin.Context)
}
