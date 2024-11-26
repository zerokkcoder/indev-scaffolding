package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zerokkcoder/indevsca/internal/logger"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// Logger GORM 自定义日志记录器
type Logger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Debug                 bool
}

// NewLogger 创建一个新的 logger 实例
func NewLogger(debug bool) *Logger {
	var slowThreshold time.Duration
	if debug {
		slowThreshold = 200 * time.Millisecond
	} else {
		slowThreshold = time.Second
	}

	return &Logger{
		SkipErrRecordNotFound: !debug,
		SlowThreshold:         slowThreshold,
		Debug:                 debug,
	}
}

// LogMode 设置日志级别
func (l *Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return l
}

// Info 记录信息
func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	logger.Info(fmt.Sprintf(msg, data...))
}

// Warn 记录警告
func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logger.Warn(fmt.Sprintf(msg, data...))
}

// Error 记录错误
func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	logger.Error(fmt.Sprintf(msg, data...))
}

// Trace 记录 SQL 语句
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	// 如果设置跳过未找到记录的错误，且错误是未找到记录，则不记录
	if l.SkipErrRecordNotFound && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	// 构建日志字段
	fields := []interface{}{
		"elapsed", elapsed,
		"rows", rows,
		"sql", sql,
	}

	if err != nil {
		fields = append(fields, "error", err)
		logger.Error("SQL执行错误", fields...)
		return
	}

	// 如果是慢查询，记录警告
	if elapsed > l.SlowThreshold {
		logger.Warn("SQL慢查询", fields...)
		return
	}

	// 在调试模式下记录所有 SQL
	if l.Debug {
		logger.Debug("SQL语句", fields...)
	}
}
