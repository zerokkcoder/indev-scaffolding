package handler

import (
	"github.com/zerokkcoder/indevsca/internal/app/handler/admin"
	"github.com/zerokkcoder/indevsca/internal/app/handler/mobile"
)

// Handlers 所有HTTP处理器的集合
type Handlers struct {
	Admin  *AdminHandlers  // 管理端处理器
	Mobile *MobileHandlers // 移动端处理器
}

// AdminHandlers 管理端处理器集合
type AdminHandlers struct {
	Auth *admin.AuthHandler // 认证相关
	User *admin.UserHandler // 用户管理
}

// MobileHandlers 移动端处理器集合
type MobileHandlers struct {
	Auth *mobile.AuthHandler // 认证相关
	User *mobile.UserHandler // 用户信息
}

// NewHandlers 创建处理器集合
func NewHandlers(
	// 管理端处理器
	adminAuth *admin.AuthHandler,
	adminUser *admin.UserHandler,
	// 移动端处理器
	mobileAuth *mobile.AuthHandler,
	mobileUser *mobile.UserHandler,
) *Handlers {
	return &Handlers{
		Admin: &AdminHandlers{
			Auth: adminAuth,
			User: adminUser,
		},
		Mobile: &MobileHandlers{
			Auth: mobileAuth,
			User: mobileUser,
		},
	}
}
