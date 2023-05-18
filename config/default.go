package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PostgresDriver string `mapstructure:"POSTGRES_DRIVER"`
	PostgresSource string `mapstructure:"POSTGRES_SOURCE"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	ClientPort     string `mapstructure:"CLIENT_PORT"`
	Origin         string `mapstructure:"ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
