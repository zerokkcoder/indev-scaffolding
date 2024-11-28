package sloggorm

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	gormlogger "gorm.io/gorm/logger"
)

const ctxLoggerKey = "slogLogger"

type Logger struct {
	SlogLogger                *slog.Logger
	SlowThreshold             time.Duration
	Colorful                  bool
	IgnoreRecordNotFoundError bool
	ParameterizedQueries      bool
	LogLevel                  gormlogger.LogLevel
}

func New(slogLogger *slog.Logger) gormlogger.Interface {
	return &Logger{
		SlogLogger:                slogLogger,
		LogLevel:                  gormlogger.Warn,
		SlowThreshold:            100 * time.Millisecond,
		Colorful:                 false,
		IgnoreRecordNotFoundError: false,
		ParameterizedQueries:      false,
	}
}

func (l *Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.logger(ctx).Info(msg, "data", fmt.Sprint(data...))
	}
}

// Warn print warn messages
func (l Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		l.logger(ctx).Warn(msg, "data", fmt.Sprint(data...))
	}
}

// Error print error messages
func (l Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		l.logger(ctx).Error(msg, "data", fmt.Sprint(data...))
	}
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	elapsedStr := fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
	logger := l.logger(ctx)

	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!errors.Is(err, gormlogger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			logger.Error("trace", 
				"error", err,
				"elapsed", elapsedStr,
				"rows", rows,
				"sql", sql,
			)
		} else {
			logger.Error("trace",
				"error", err,
				"elapsed", elapsedStr,
				"rows", rows,
				"sql", sql,
			)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			logger.Warn("trace",
				"slow", slowLog,
				"elapsed", elapsedStr,
				"rows", rows,
				"sql", sql,
			)
		} else {
			logger.Warn("trace",
				"slow", slowLog,
				"elapsed", elapsedStr,
				"rows", rows,
				"sql", sql,
			)
		}
	case l.LogLevel == gormlogger.Info:
		sql, rows := fc()
		if rows == -1 {
			logger.Info("trace",
				"elapsed", elapsedStr,
				"rows", rows,
				"sql", sql,
			)
		} else {
			logger.Info("trace",
				"elapsed", elapsedStr,
				"rows", rows,
				"sql", sql,
			)
		}
	}
}

var (
	gormPackage = filepath.Join("gorm.io", "gorm")
)

func (l Logger) logger(ctx context.Context) *slog.Logger {
	logger := l.SlogLogger
	if ctx != nil {
		if c, ok := ctx.(*gin.Context); ok {
			ctx = c.Request.Context()
		}
		sl := ctx.Value(ctxLoggerKey)
		ctxLogger, ok := sl.(*slog.Logger)
		if ok {
			logger = ctxLogger
		}
	}

	// 添加调用者信息
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		default:
			return logger.With(
				"caller", fmt.Sprintf("%s:%d", file, line),
			)
		}
	}
	return logger
}
