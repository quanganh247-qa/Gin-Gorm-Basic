package util

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Host                string        `mapstructure:"HOST"`
	DSN                 string        `mapstructure:"DSN"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	SymmetricKey        string        `mapstructure:"SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	ApiPrefix           string        `mapstructure:"API_PREFIX"`
}

var Configs = Config{}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		LoadConfig(path)
	})

	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&Configs)
	return &Configs, err
}
