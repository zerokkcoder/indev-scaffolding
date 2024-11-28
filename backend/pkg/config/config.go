package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 配置
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Log      LogConfig      `mapstructure:"log"`
}

// AppConfig 应用配置
type AppConfig struct {
	Name      string `mapstructure:"name"`
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Mode      string `mapstructure:"mode"`
	Debug     bool   `mapstructure:"debug"`
	JWTSecret string `mapstructure:"jwt_secret"`
	Cache     string `mapstructure:"cache"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Charset  string `mapstructure:"charset"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level  string `mapstructure:"level"`
	Output string `mapstructure:"output"`
	Dir    string `mapstructure:"dir"`
}

// NewConfig 创建配置
func NewConfig(path string) (*Config, error) {
	// 设置默认path
	if path == "" {
		path = "config/local.yml"
	}
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var conf Config
	if err := v.Unmarshal(&conf); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &conf, nil
}
