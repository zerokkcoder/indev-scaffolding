package config

import "github.com/spf13/viper"

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Output     string `mapstructure:"output"`
	Dir        string `mapstructure:"dir"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

func setLogDefaults() {
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.output", "stdout") // stdout, file
	viper.SetDefault("log.dir", "./storages/logs")
	viper.SetDefault("log.max_size", 100) // 100MB
	viper.SetDefault("log.max_backups", 10)
	viper.SetDefault("log.max_age", 30)    // 30天
	viper.SetDefault("log.compress", true) // 是否压缩
}
