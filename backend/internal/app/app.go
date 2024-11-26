package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/app/handler"
	"github.com/zerokkcoder/indevsca/internal/app/routes"
	"github.com/zerokkcoder/indevsca/internal/infra/config"
	"gorm.io/gorm"
	"github.com/zerokkcoder/indevsca/internal/infra/seed"
)

// App 应用程序
type App struct {
	config   *config.Config
	engine   *gin.Engine
	handlers *handler.Handlers
	server   *http.Server
	db       *gorm.DB
}

// NewApp 创建应用程序
func NewApp(
	cfg *config.Config,
	handlers *handler.Handlers,
	db *gorm.DB,
) *App {
	app := &App{
		config:   cfg,
		handlers: handlers,
		engine:   gin.Default(),
		db:       db,
	}

	gin.SetMode(cfg.App.Mode)
	app.setupRoutes()
	app.setupServer()
	app.seedData() // 添加数据种子

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
	api := a.engine.Group("/api")
	routes.RegisterAdminRoutes(api, a.handlers.Admin)
	routes.RegisterMobileRoutes(api, a.handlers.Mobile)
}

// seedData 初始化数据
func (a *App) seedData() error {
	return seed.CreateSuperAdmin(a.db)
}
