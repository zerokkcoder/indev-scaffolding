package mobile

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/domain/service"
)

// AuthHandler 移动端认证处理器
type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler 创建移动端认证处理器
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login 用户登录
func (h *AuthHandler) Login(c *gin.Context) {
	// 实现移动端用户登录逻辑
}

// Register 用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	// 实现移动端用户注册逻辑
}

// SendVerificationCode 发送验证码
func (h *AuthHandler) SendVerificationCode(c *gin.Context) {
	// 实现发送验证码逻辑
}
