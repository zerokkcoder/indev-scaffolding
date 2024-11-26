package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Driver          string        `mapstructure:"driver"`            // 数据库类型：sqlite, postgres, mysql
	Host            string        `mapstructure:"host"`              // 数据库主机
	Port            int           `mapstructure:"port"`              // 数据库端口
	Database        string        `mapstructure:"database"`          // 数据库名
	User            string        `mapstructure:"user"`              // 用户名
	Password        string        `mapstructure:"password"`          // 密码
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`    // 最大空闲连接数
	MaxOpenConns    int           `mapstructure:"max_open_conns"`    // 最大打开连接数
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"` // 连接最大生命周期（秒）
	LogLevel        string        `mapstructure:"log_level"`         // 日志级别
}

// GetDSN 根据数据库类型返回对应的 DSN
func (c *DatabaseConfig) GetDSN() string {
	switch c.Driver {
	case "sqlite":
		return fmt.Sprintf("./storages/data/%s.db", c.Database)
	case "postgres":
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Password, c.Database)
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.User, c.Password, c.Host, c.Port, c.Database)
	default:
		return ""
	}
}

func setDatabaseDefaults() {
	viper.SetDefault("database.driver", "postgres")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "postgres")
	viper.SetDefault("database.database", "indever")
	viper.SetDefault("database.max_open_conns", 100)
	viper.SetDefault("database.max_idle_conns", 10)
	viper.SetDefault("database.conn_max_lifetime", 5*time.Minute)
	viper.SetDefault("database.log_level", "info")
}
