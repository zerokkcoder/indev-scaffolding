package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/domain/service"
)

// AuthHandler 管理员认证处理器
type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler 创建管理员认证处理器
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login 管理员登录
func (h *AuthHandler) Login(c *gin.Context) {
	// 实现管理员登录逻辑
}

// Logout 管理员登出
func (h *AuthHandler) Logout(c *gin.Context) {
	// 实现管理员登出逻辑
}
