package management

import (
	"sync"
	"time"

	"github.com/spf13/viper"
)

type (
	DeviceSettings struct {
		Type string
	}

	ConnectionSettings struct {
		ServerHost string
		ConnTimeout time.Duration
		MaxIdleConns int
		IdleConnTimeout time.Duration
		DisableCompression bool
	}

	DeviceConfig struct {
		DeviceSettings *DeviceSettings
		ConnectionSettings *ConnectionSettings
	}
)

var (
	once sync.Once
	configInstance *DeviceConfig
	configError error
)

func GetConfig() (*DeviceConfig, error) {
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