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
	IdleTimeout  int    `mapstructure:"idle_timeout"`
	MaxRetries   int    `mapstructure:"max_retries"`
	RetryBackoff int    `mapstructure:"retry_backoff"`
}

// GetAddr 获取Redis地址
func (c *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func setRedisDefaults() {
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("redis.pool_size", 100)
	viper.SetDefault("redis.min_idle_conns", 10)
	viper.SetDefault("redis.idle_timeout", 5*time.Minute)
	viper.SetDefault("redis.max_retries", 3)
	viper.SetDefault("redis.retry_backoff", 100*time.Millisecond)
}
