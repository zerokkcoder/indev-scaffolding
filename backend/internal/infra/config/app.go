package config

import "github.com/spf13/viper"

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

func setAppDefaults() {
	viper.SetDefault("app.name", "Indever")
	viper.SetDefault("app.host", "0.0.0.0")
	viper.SetDefault("app.port", 8088)
	viper.SetDefault("app.mode", "debug")
	viper.SetDefault("app.debug", false)
	viper.SetDefault("app.jwt_secret", "")
	viper.SetDefault("app.cache", "redis")
}

// GetPort 获取端口号，默认8080
func (c *AppConfig) GetPort() int {
	if c.Port == 0 {
		return 8080
	}
	return c.Port
}

// GetMode 获取运行模式，默认debug
func (c *AppConfig) GetMode() string {
	if c.Mode == "" {
		return "debug"
	}
	return c.Mode
}
