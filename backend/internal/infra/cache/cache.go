package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zerokkcoder/indevsca/internal/infra/config"
	"github.com/zerokkcoder/indevsca/internal/infra/logger"
)

var (
	instance *Cache
	once     sync.Once
)

// Cache Redis缓存管理器
type Cache struct {
	client *redis.Client
}

// New 创建缓存实例
func New(cfg *config.Config) (*Cache, error) {
	if cfg.App.Cache != "redis" {
		return nil, nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.GetAddr(),
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     cfg.Redis.GetPoolSize(),
		MinIdleConns: cfg.Redis.GetMinIdleConns(),
	})

	// 测试连接
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &Cache{client: client}, nil
}

// GetInstance 获取缓存实例（单例模式）
func GetInstance() *Cache {
	once.Do(func() {
		instance = &Cache{}
		if err := instance.init(); err != nil {
			logger.Error("初始化缓存失败", "error", err)
			panic(err)
		}
	})
	return instance
}

// init 初始化Redis连接
func (c *Cache) init() error {
	// 是否使用缓存
	if config.Get().App.Cache != "redis" {
		return nil
	}
	cfg := config.Get().Redis
	c.client = redis.NewClient(&redis.Options{
		Addr:         cfg.GetAddr(),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	// 测试连接
	ctx := context.Background()
	if err := c.client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("连接Redis失败: %w", err)
	}

	return nil
}

// Set 设置缓存
func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

// Get 获取缓存
func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// Delete 删除缓存
func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Close 关闭连接
func (c *Cache) Close() error {
	return c.client.Close()
}
