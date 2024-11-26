package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/app/handler"
)

// RegisterMobileRoutes 注册移动端路由
func RegisterMobileRoutes(api *gin.RouterGroup, handlers *handler.MobileHandlers) {
	mobile := api.Group("/mobile")
	{
		// 认证路由
		auth := mobile.Group("/auth")
		{
			auth.POST("/login", handlers.Auth.Login)
			auth.POST("/register", handlers.Auth.Register)
			auth.POST("/verify-code", handlers.Auth.SendVerificationCode)
		}

		// 用户信息路由
		user := mobile.Group("/user")
		{
			user.GET("/profile", handlers.User.Profile)           // 获取用户信息
			user.PUT("/profile", handlers.User.UpdateProfile)     // 更新用户信息
			user.PUT("/password", handlers.User.UpdatePassword)   // 修改密码
			user.POST("/bind/phone", handlers.User.BindPhone)     // 绑定手机号
			user.POST("/bind/email", handlers.User.BindEmail)     // 绑定邮箱
		}
	}
}
