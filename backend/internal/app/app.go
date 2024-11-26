package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/app/handler"
	"github.com/zerokkcoder/indevsca/internal/infra/config"
)

// App 应用程序
type App struct {
	config   *config.Config
	engine   *gin.Engine
	handlers *handler.Handlers
	server   *http.Server
}

// NewApp 创建应用程序
func NewApp(
	cfg *config.Config,
	handlers *handler.Handlers,
) *App {
    app := &App{
        config:   cfg,
        handlers: handlers,
        engine:   gin.Default(),
    }

    gin.SetMode(cfg.App.Mode)
    app.setupRoutes()
    app.setupServer()

    return app
}

// setupServer 设置HTTP服务器
func (a *App) setupServer() {
	a.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.App.Port),
		Handler: a.engine,
	}
}

// Start 启动应用
func (a *App) Start() error {
	return a.server.ListenAndServe()
}

// Stop 优雅关闭
func (a *App) Stop(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

// setupRoutes 设置路由
func (a *App) setupRoutes() {
	// 健康检查
	a.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// API路由组
	api := a.engine.Group("/api")
	{
		// 认证路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", a.handlers.Auth.Register)
			auth.POST("/login", a.handlers.Auth.Login)
		}
	}
}
