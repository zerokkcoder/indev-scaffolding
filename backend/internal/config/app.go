package config

import "github.com/spf13/viper"

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Host      string `mapstructure:"host"`
	Debug     bool   `mapstructure:"debug"`
	Port      int    `mapstructure:"port"`
	JWTSecret string `mapstructure:"jwt_secret"`
	UploadDir string `mapstructure:"upload_dir"`
}

func setAppDefaults() {
	viper.SetDefault("app.name", "Indever")
	viper.SetDefault("app.host", "0.0.0.0")
	viper.SetDefault("app.port", 8088)
	viper.SetDefault("app.debug", false)
	viper.SetDefault("app.jwt_secret", "")
	viper.SetDefault("app.upload_dir", "./storages/uploads")
}
