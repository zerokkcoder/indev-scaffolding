package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/app/handler"
	"github.com/zerokkcoder/indevsca/internal/app/routes"
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
	routes.RegisterRoutes(a.engine, a.handlers)
}
