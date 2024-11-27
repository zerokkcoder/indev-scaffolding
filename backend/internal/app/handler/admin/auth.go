package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/app/requests/admin"
	adminService "github.com/zerokkcoder/indevsca/internal/domain/service/admin"
	"github.com/zerokkcoder/indevsca/internal/pkg/response"
	"gorm.io/gorm"
)

// AuthHandler 管理员认证处理器
type AuthHandler struct {
	authService *adminService.AuthService
}

// NewAuthHandler 创建管理员认证处理器
func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		authService: adminService.NewAuthService(db),
	}
}

// Login 管理员登录
func (h *AuthHandler) Login(c *gin.Context) {
	var req admin.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ErrInvalidParams)
		return
	}

	token, admin, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, gin.H{
		"token": token,
		"admin": admin,
	})
}

// Logout 管理员登出
func (h *AuthHandler) Logout(c *gin.Context) {
	// 清除认证相关的 header
	c.Header("Authorization", "")

	response.Success(c, gin.H{
		"message": "退出成功",
	})
}
