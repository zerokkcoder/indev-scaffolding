package log

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/zerokkcoder/indevsca/pkg/config"
)

// Logger 日志记录器
type Logger struct {
	*slog.Logger
	writer io.Writer
}

// NewLogger 创建日志记录器
func NewLogger(conf *config.Config) *Logger {
	var writer io.Writer

	// 在debug模式下同时输出到控制台和文件
	if conf.App.Mode == "debug" {
		if conf.Log.Output == "file" {
			fileWriter := createFileWriter(conf.Log.Dir)
			writer = io.MultiWriter(os.Stdout, fileWriter)
		} else {
			writer = os.Stdout
		}
	} else {
		switch conf.Log.Output {
		case "file":
			writer = createFileWriter(conf.Log.Dir)
		case "stdout":
			writer = os.Stdout
		default:
			writer = os.Stdout
		}
	}

	// 设置日志选项
	opts := &slog.HandlerOptions{
		Level:     parseLevel(conf.Log.Level),
		AddSource: true,
	}

	// 创建JSON处理器
	handler := slog.NewJSONHandler(writer, opts)

	// 创建logger
	logger := slog.New(handler)

	// 设置默认logger
	slog.SetDefault(logger)

	return &Logger{
		Logger: logger,
		writer: writer,
	}
}

// parseLevel 解析日志级别
func parseLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// createFileWriter 创建文件写入器
func createFileWriter(dir string) io.Writer {
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	logFile := filepath.Join(dir, time.Now().Format("2006-01-02")+".log")
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("创建日志文件失败: %v", err))
	}

	return f
}

// Debug 输出Debug级别日志
func (l *Logger) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

// Info 输出Info级别日志
func (l *Logger) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

// Warn 输出Warn级别日志
func (l *Logger) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

// Error 输出Error级别日志
func (l *Logger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

// Close 关闭日志文件
func (l *Logger) Close() error {
	if closer, ok := l.writer.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
