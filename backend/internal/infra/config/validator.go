package config

import "fmt"

// Validate 验证配置
func (c *Config) Validate() error {
	if err := c.App.Validate(); err != nil {
		return fmt.Errorf("app config: %w", err)
	}
	if err := c.Database.Validate(); err != nil {
		return fmt.Errorf("database config: %w", err)
	}
	if err := c.Redis.Validate(); err != nil {
		return fmt.Errorf("redis config: %w", err)
	}
	if err := c.Log.Validate(); err != nil {
		return fmt.Errorf("log config: %w", err)
	}
	return nil
}

// Validate 验证应用配置
func (c *AppConfig) Validate() error {
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("invalid port: %d", c.Port)
	}
	if c.Mode != "debug" && c.Mode != "release" && c.Mode != "test" {
		return fmt.Errorf("invalid mode: %s", c.Mode)
	}
	return nil
}

// Validate 验证数据库配置
func (c *DatabaseConfig) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("invalid port: %d", c.Port)
	}
	if c.Database == "" {
		return fmt.Errorf("database is required")
	}
	if c.Username == "" {
		return fmt.Errorf("username is required")
	}
	return nil
}

// Validate 验证Redis配置
func (c *RedisConfig) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("invalid port: %d", c.Port)
	}
	if c.DB < 0 {
		return fmt.Errorf("invalid db: %d", c.DB)
	}
	return nil
}

// Validate 验证日志配置
func (c *LogConfig) Validate() error {
	if c.Level == "" {
		return fmt.Errorf("level is required")
	}
	if c.Output == "" {
		return fmt.Errorf("output is required")
	}
	return nil
}
