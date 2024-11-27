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
	if _, err := config.LoadConfig(); err != nil {
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
		return // 确保在初始化失败时退出
	}

	// 启动应用
	errChan := make(chan error, 1)
	go func() {
		logger.Info("Starting server")
		if err := application.Start(); err != nil {
			errChan <- err
		}
	}()

	// 等待中断信号或错误
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		logger.Error("Server failed", "error", err)
	case <-quit:
		logger.Info("Shutting down server...")

		// 优雅关闭服务器（5秒超时）
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := application.Stop(ctx); err != nil {
			logger.Error("Server forced to shutdown", "error", err)
		}

		logger.Info("Server exited")
	}
}
