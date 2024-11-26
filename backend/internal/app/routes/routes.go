package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/app/handler"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(engine *gin.Engine, handlers *handler.Handlers) {
	// 健康检查
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// API路由组
	api := engine.Group("/api")
	{
		// 注册管理端路由
		RegisterAdminRoutes(api, handlers.Admin)

		// 注册移动端路由
		RegisterMobileRoutes(api, handlers.Mobile)
	}
}
