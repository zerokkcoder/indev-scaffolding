package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/domain/entity"
	"github.com/zerokkcoder/indevsca/internal/domain/service"
	"github.com/zerokkcoder/indevsca/internal/pkg/response"
	"github.com/zerokkcoder/indevsca/internal/pkg/utils"
	"gorm.io/gorm"
)

// AuthHandler 管理员认证处理器
type AuthHandler struct {
	authService *service.AuthService
	db          *gorm.DB
}

// NewAuthHandler 创建管理员认证处理器
func NewAuthHandler(authService *service.AuthService, db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		db:          db,
	}
}

// LoginRequest 管理员登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 管理员登录
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ErrInvalidParams)
		return
	}

	var admin entity.Admin
	if err := h.db.Where("username = ?", req.Username).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Error(c, response.ErrInvalidCredentials)
			return
		}
		response.Error(c, err)
		return
	}

	// 验证密码
	if !utils.CheckPasswordHash(req.Password, admin.Password) {
		response.Error(c, response.ErrInvalidCredentials)
		return
	}

	// 生成JWT token
	token, err := utils.GenerateToken(admin.ID)
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
