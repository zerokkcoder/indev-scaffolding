package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/app/handler"
)

// RegisterAdminRoutes 注册管理端路由
func RegisterAdminRoutes(api *gin.RouterGroup, handlers *handler.AdminHandlers) {
	admin := api.Group("/admin")
	{
		// 认证路由
		auth := admin.Group("/auth")
		{
			auth.POST("/login", handlers.Auth.Login)
			auth.POST("/logout", handlers.Auth.Logout)
		}

		// 用户管理路由
		users := admin.Group("/users")
		{
			users.GET("", handlers.User.List)                 // 获取用户列表
			users.PUT("/:id", handlers.User.Update)           // 更新用户
			users.DELETE("/:id", handlers.User.Delete)        // 删除用户
			users.POST("/:id/enable", handlers.User.Enable)   // 启用用户
			users.POST("/:id/disable", handlers.User.Disable) // 禁用用户
		}
	}
}
