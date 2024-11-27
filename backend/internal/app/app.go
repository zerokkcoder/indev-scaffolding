package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/app/handler"
	"github.com/zerokkcoder/indevsca/internal/app/middleware"
	"github.com/zerokkcoder/indevsca/internal/app/routes"
	"github.com/zerokkcoder/indevsca/internal/infra/config"
	"github.com/zerokkcoder/indevsca/internal/pkg/validator"
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
	if cfg == nil {
		panic("config is required")
	}
	if handlers == nil {
		panic("handlers is required")
	}
	if db == nil {
		panic("database is required")
	}

	// 验证配置
	if err := cfg.Validate(); err != nil {
		panic(fmt.Sprintf("invalid config: %v", err))
	}

	app := &App{
		config:   cfg,
		handlers: handlers,
		engine:   gin.New(),
		db:       db,
	}

	// 设置运行模式
	gin.SetMode(cfg.App.Mode)

	// 初始化验证器
	validator.Setup()

	app.setupMiddlewares()
	app.setupRoutes()
	app.setupServer()
	if err := app.seedData(); err != nil {
		panic(fmt.Sprintf("failed to seed data: %v", err))
	}

	return app
}

// setupServer 设置HTTP服务器
func (a *App) setupServer() {
	if a.config == nil {
		panic("config is not initialized")
	}
	if a.engine == nil {
		panic("engine is not initialized")
	}

	a.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.App.Port),
		Handler: a.engine,
	}
}

// Start 启动应用
func (a *App) Start() error {
	if a.server == nil {
		return fmt.Errorf("server is not initialized")
	}
	return a.server.ListenAndServe()
}

// Stop 优雅关闭
func (a *App) Stop(ctx context.Context) error {
	if a.server == nil {
		return nil
	}
	return a.server.Shutdown(ctx)
}

// setupRoutes 设置路由
func (a *App) setupRoutes() {
	if a.engine == nil {
		panic("engine is not initialized")
	}
	if a.handlers == nil {
		panic("handlers is not initialized")
	}

	api := a.engine.Group("/api")
	routes.RegisterAdminRoutes(api, a.handlers.Admin)
	routes.RegisterMobileRoutes(api, a.handlers.Mobile)
}

// setupMiddlewares 设置中间件
func (a *App) setupMiddlewares() {
	if a.engine == nil {
		panic("engine is not initialized")
	}

	// 全局中间件
	a.engine.Use(middleware.Logger())    // 日志中间件
	a.engine.Use(middleware.CORS())      // 跨域中间件
	a.engine.Use(gin.Recovery())         // panic恢复中间件
}

// seedData 初始化数据
func (a *App) seedData() error {
	if a.db == nil {
		return fmt.Errorf("database is not initialized")
	}
	return seed.CreateSuperAdmin(a.db)
}
