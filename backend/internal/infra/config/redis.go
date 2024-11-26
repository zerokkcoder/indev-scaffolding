package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

// GetAddr 获取Redis地址
func (c *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// GetPoolSize 获取连接池大小，默认10
func (c *RedisConfig) GetPoolSize() int {
	if c.PoolSize == 0 {
		return 10
	}
	return c.PoolSize
}

// GetMinIdleConns 获取最小空闲连接数，默认2
func (c *RedisConfig) GetMinIdleConns() int {
	if c.MinIdleConns == 0 {
		return 2
	}
	return c.MinIdleConns
}

// GetConnParams 获取Redis连接参数
func (c *RedisConfig) GetConnParams() map[string]interface{} {
	return map[string]interface{}{
		"addr":         c.GetAddr(),
		"password":     c.Password,
		"db":           c.DB,
		"poolSize":     c.GetPoolSize(),
		"minIdleConns": c.GetMinIdleConns(),
		"dialTimeout":  5 * time.Second,
		"readTimeout":  3 * time.Second,
		"writeTimeout": 3 * time.Second,
		"poolTimeout":  4 * time.Second,
	}
}

func setRedisDefaults() {
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("redis.pool_size", 100)
	viper.SetDefault("redis.min_idle_conns", 10)
}
