package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Log      LogConfig      `mapstructure:"log"`
}

var (
	globalConfig *Config
)

// Init 初始化配置
func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 设置所有配置的默认值
	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置到结构体
	globalConfig = &Config{}
	if err := viper.Unmarshal(globalConfig); err != nil {
		return fmt.Errorf("解析配置失败: %w", err)
	}

	// 确保必要的目录存在
	if err := ensureDirectories(); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	return nil
}

// Get 获取全局配置
func Get() *Config {
	return globalConfig
}

func setDefaults() {
	setAppDefaults()
	setDatabaseDefaults()
	setRedisDefaults()
	setLogDefaults()
}

func ensureDirectories() error {
	dirs := []string{
		globalConfig.App.UploadDir,
		globalConfig.Log.Dir,
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("创建目录 %s 失败: %w", dir, err)
		}
	}
	return nil
}
