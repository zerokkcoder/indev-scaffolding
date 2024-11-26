package database

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/zerokkcoder/indevsca/internal/config"
	"github.com/zerokkcoder/indevsca/internal/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	instance *Database
	once     sync.Once
)

// Database 数据库管理器
type Database struct {
	db *gorm.DB
}

// GetInstance 获取数据库实例（单例模式）
func GetInstance() *Database {
	once.Do(func() {
		instance = &Database{}
		if err := instance.init(); err != nil {
			logger.Error("初始化数据库失败", "error", err)
			panic(err)
		}
	})
	return instance
}

// init 初始化数据库连接
func (d *Database) init() error {
	cfg := config.Get()

	// 创建 GORM 配置
	gormConfig := &gorm.Config{
		Logger: NewLogger(cfg.App.Debug),
	}

	// 在 debug 模式下启用调试
	if cfg.App.Debug {
		gormConfig.PrepareStmt = true // 缓存预编译语句
		gormConfig.QueryFields = true // 显式指定查询字段
	}

	var err error
	dsn := cfg.Database.GetDSN()
	switch cfg.Database.Driver {
	case "sqlite":
		// 确保数据库文件目录存在
		if err := os.MkdirAll(filepath.Dir(dsn), 0755); err != nil {
			return fmt.Errorf("创建数据库目录失败: %w", err)
		}
		d.db, err = gorm.Open(sqlite.Open(dsn), gormConfig)

	case "postgres":
		d.db, err = gorm.Open(postgres.Open(dsn), gormConfig)

	case "mysql":
		d.db, err = gorm.Open(mysql.Open(dsn), gormConfig)

	default:
		return fmt.Errorf("不支持的数据库类型: %s", cfg.Database.Driver)
	}

	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	// 配置连接池
	sqlDB, err := d.db.DB()
	if err != nil {
		return fmt.Errorf("获取数据库实例失败: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * cfg.Database.ConnMaxLifetime)

	// Debug 模式下开启 SQL 调试
	if cfg.App.Debug {
		d.db = d.db.Debug()
	}

	return nil
}

// DB 获取 GORM 数据库实例
func (d *Database) DB() *gorm.DB {
	return d.db
}

// Transaction 执行事务
func (d *Database) Transaction(fn func(tx *gorm.DB) error) error {
	return d.db.Transaction(fn)
}

// AutoMigrate 自动迁移数据库表结构
func (d *Database) AutoMigrate(models ...interface{}) error {
	return d.db.AutoMigrate(models...)
}

// Close 关闭数据库连接
func (d *Database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
