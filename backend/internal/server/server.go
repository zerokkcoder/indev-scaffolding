package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/indevsca/internal/config"
	"github.com/zerokkcoder/indevsca/internal/database"
	"github.com/zerokkcoder/indevsca/internal/logger"
	"github.com/zerokkcoder/indevsca/internal/middlewares"
)

var (
	instance *Server
	once     sync.Once
)

// Server 服务器实例
type Server struct {
	engine *gin.Engine
	srv    *http.Server
	db     *database.Database
}

// GetInstance 获取单例
func GetInstance() *Server {
	once.Do(func() {
		instance = &Server{}
		instance.init()
	})
	return instance
}

// init 初始化服务器
func (s *Server) init() {
	cfg := config.Get()

	// 初始化数据库
	s.db = database.GetInstance()

	// 运行数据库迁移
	migrator := database.NewMigrator(s.db.DB())
	if err := migrator.RunMigrations(); err != nil {
		logger.Error("数据库迁移失败", "error", err)
		panic(err)
	}

	// 设置 gin 模式
	if cfg.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化 gin 引擎
	s.engine = gin.New()
	// 中间件
	s.engine.Use(middlewares.Logger(), middlewares.Recovery())

	// 初始化 http 服务
	s.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port),
		Handler: s.engine,
	}

	// TODO: 初始化路由
	s.setupRoutes()
}

// Start 启动服务器
func (s *Server) Start() error {
	cfg := config.Get()
	logger.Info("启动服务器", "host", cfg.App.Host, "port", cfg.App.Port)
	return s.srv.ListenAndServe()
}

// Stop 优雅关闭服务器
func (s *Server) Stop(ctx context.Context) error {
	logger.Info("正在关闭服务器...")
	return s.srv.Shutdown(ctx)
}

// setupRoutes 设置路由
func (s *Server) setupRoutes() {
	// TODO: 添加路由配置
	// 测试路由
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
