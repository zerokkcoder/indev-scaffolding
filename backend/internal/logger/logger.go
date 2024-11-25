package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/zerokkcoder/indevsca/internal/config"
	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var defaultLogger *slog.Logger

// Options 日志选项
type Options struct {
	Level      slog.Level
	Output     string
	Dir        string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

// multiWriter 多路输出
type multiWriter struct {
	writers []io.Writer
}

func (t *multiWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
	}
	return len(p), nil
}

// Init 初始化日志
func Init() error {
	cfg := config.Get().Log

	level := parseLevel(cfg.Level)
	opts := &Options{
		Level:      level,
		Output:     cfg.Output,
		Dir:        cfg.Dir,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	handler, err := newHandler(opts)
	if err != nil {
		return fmt.Errorf("创建日志处理器失败: %w", err)
	}

	defaultLogger = slog.New(handler)
	slog.SetDefault(defaultLogger)

	return nil
}

// Debug 输出调试日志
func Debug(msg string, args ...any) {
	defaultLogger.Debug(msg, args...)
}

// Info 输出信息日志
func Info(msg string, args ...any) {
	defaultLogger.Info(msg, args...)
}

// Warn 输出警告日志
func Warn(msg string, args ...any) {
	defaultLogger.Warn(msg, args...)
}

// Error 输出错误日志
func Error(msg string, args ...any) {
	defaultLogger.Error(msg, args...)
}

// With 返回带有额外属性的日志记录器
func With(args ...any) *slog.Logger {
	return defaultLogger.With(args...)
}

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

func newHandler(opts *Options) (slog.Handler, error) {
	var writer io.Writer
	writers := []io.Writer{}

	// 在开发环境下总是添加标准输出
	if config.Get().App.Debug {
		writers = append(writers, os.Stdout)
	}

	// 如果配置了文件输出，添加文件writer
	if opts.Output == "file" {
		// 使用 lumberjack 进行日志切割
		fileWriter := &lumberjack.Logger{
			Filename:   filepath.Join(opts.Dir, fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))),
			MaxSize:    opts.MaxSize, // MB
			MaxBackups: opts.MaxBackups,
			MaxAge:     opts.MaxAge, // days
			Compress:   opts.Compress,
		}
		writers = append(writers, fileWriter)
	} else {
		// 非开发环境且未配置文件输出时，使用标准输出
		writers = append(writers, os.Stdout)
	}
	
	// 根据writer数量选择输出方式
	switch len(writers) {
	case 0:
		writer = os.Stdout // 保底输出
	case 1:
		writer = writers[0]
	default:
		writer = &multiWriter{writers: writers}
	}
	
	return slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: opts.Level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// 添加时间戳
			if a.Key == slog.TimeKey {
				return slog.Attr{
					Key:   "timestamp",
					Value: slog.StringValue(a.Value.Time().Format(time.RFC3339)),
				}
			}
			return a
		},
	}), nil
}
