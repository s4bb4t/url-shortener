package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Stage      string   `mapstructure:"stage" yaml:"stage"`
	HTTPServer Server   `mapstructure:"http_server" yaml:"http_server"`
	Postgres   Postgres `mapstructure:"postgres" yaml:"postgres"`
}

func NewConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return nil, errors.New("CONFIG_PATH environment variable not set")
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %s", err)
	}
	return &cfg, nil
}
