package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Driver        string `mapstructure:"driver"`         // 数据库类型：sqlite, postgres, mysql
	Host          string `mapstructure:"host"`           // 数据库主机
	Port          int    `mapstructure:"port"`           // 数据库端口
	Username      string `mapstructure:"username"`       // 数据库用户名
	Password      string `mapstructure:"password"`       // 数据库密码
	Database      string `mapstructure:"database"`       // 数据库名
	Charset       string `mapstructure:"charset"`        // 字符集
	MaxIdleConns  int    `mapstructure:"max_idle_conns"` // 最大空闲连接数
	MaxOpenConns  int    `mapstructure:"max_open_conns"` // 最大打开连接数
	MaxLifetime   int    `mapstructure:"max_lifetime"`   // 连接最大生命周期（秒）
	SlowThreshold int    `mapstructure:"slow_threshold"` // 慢查询阈值（毫秒）
}

// GetDSN 根据数据库类型返回对应的 DSN
func (c *DatabaseConfig) GetDSN() string {
	switch c.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			c.Username, c.Password, c.Host, c.Port, c.Database, c.GetCharset())
	default:
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			c.Username, c.Password, c.Host, c.Port, c.Database, c.GetCharset())
	}
}

// GetCharset 获取字符集，默认utf8mb4
func (c *DatabaseConfig) GetCharset() string {
	if c.Charset == "" {
		return "utf8mb4"
	}
	return c.Charset
}

// GetMaxIdleConns 获取最大空闲连接数，默认10
func (c *DatabaseConfig) GetMaxIdleConns() int {
	if c.MaxIdleConns == 0 {
		return 10
	}
	return c.MaxIdleConns
}

// GetMaxOpenConns 获取最大打开连接数，默认100
func (c *DatabaseConfig) GetMaxOpenConns() int {
	if c.MaxOpenConns == 0 {
		return 100
	}
	return c.MaxOpenConns
}

// GetMaxLifetime 获取连接最大生命周期，默认3600秒
func (c *DatabaseConfig) GetMaxLifetime() int {
	if c.MaxLifetime == 0 {
		return 3600
	}
	return c.MaxLifetime
}

// GetSlowThreshold 获取慢查询阈值，默认200毫秒
func (c *DatabaseConfig) GetSlowThreshold() int {
	if c.SlowThreshold == 0 {
		return 200
	}
	return c.SlowThreshold
}

func setDatabaseDefaults() {
	viper.SetDefault("database.driver", "mysql")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.username", "")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.database", "indever")
	viper.SetDefault("database.charset", "utf8mb4")
	viper.SetDefault("database.max_open_conns", 100)
	viper.SetDefault("database.max_idle_conns", 10)
	viper.SetDefault("database.conn_max_lifetime", 5*time.Minute)
	viper.SetDefault("database.slow_threshold", 800)
}
