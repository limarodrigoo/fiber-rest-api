package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseUser     string `mapstructure:"DB_USER"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabasePort     string `mapstructure:"DB_PORT"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	return
}
