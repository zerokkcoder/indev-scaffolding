package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zerokkcoder/indevsca/internal/app"
	"github.com/zerokkcoder/indevsca/internal/infra/config"
	"github.com/zerokkcoder/indevsca/internal/infra/logger"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// 初始化应用
	application, err := app.InitializeApp()
	if err != nil {
		logger.Fatal("Failed to initialize application", "error", err)
	}

	// 启动应用
	go func() {
		logger.Info("Starting server", "port", cfg.App.Port)
		if err := application.Start(); err != nil {
			logger.Fatal("Failed to start server", "error", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器,为关闭服务器操作设置一个5秒的超时
	// 创建一个接收信号的通道
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")
	// 优雅关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := application.Stop(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
	}

	logger.Info("Server exited")
}
