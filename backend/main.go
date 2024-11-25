package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zerokkcoder/indevsca/internal/config"
	"github.com/zerokkcoder/indevsca/internal/logger"
	"github.com/zerokkcoder/indevsca/internal/server"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		panic("初始化配置失败: " + err.Error())
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		panic("初始化日志失败: " + err.Error())
	}

	// 获取服务器实例
	srv := server.GetInstance()

	// 在后台启动服务器
	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			logger.Error("启动服务器失败", "error", err)
			os.Exit(1)
		}
	}()

	// 等待中断信号来优雅地关闭服务器,为关闭服务器操作设置一个5秒的超时
	// 创建一个接收信号的通道
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 优雅关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		logger.Error("优雅关闭服务器失败", "error", err)
	}

	logger.Info("服务器已优雅关闭")
}
