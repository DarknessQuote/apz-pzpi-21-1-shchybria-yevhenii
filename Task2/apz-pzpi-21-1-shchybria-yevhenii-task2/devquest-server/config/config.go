package config

import (
	"sync"

	"github.com/spf13/viper"
)

type (
	Server struct {
		Port int
	}

	Config struct {
		Server *Server
	}
)

var (
	once sync.Once
	configInstance *Config
	configError error
)

func GetConfig() (*Config, error) {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")

		if err := viper.ReadInConfig(); err != nil {
			configError = err
			return
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			configError = err
			return
		}
	})

	return configInstance, configError
}