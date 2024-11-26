package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	globalConfig *Config
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Log      LogConfig      `mapstructure:"log"`
}

// New 创建配置实例
func New() (*Config, error) {
	return LoadConfig()
}

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	if globalConfig != nil {
		return globalConfig, nil
	}

	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	// 环境变量配置
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// 设置默认值
	setDefaults()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// 解析配置到结构体
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	globalConfig = &cfg
	return globalConfig, nil
}

// Get 获取全局配置
func Get() *Config {
	if globalConfig == nil {
		panic("config not loaded")
	}
	return globalConfig
}

// setDefaults 设置默认值
func setDefaults() {
	setAppDefaults()
	setDatabaseDefaults()
	setRedisDefaults()
	setLogDefaults()
}
