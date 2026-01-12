package config

import "time"

type Server struct {
	Host        string        `yaml:"host" mapstructure:"host"`
	Port        string        `yaml:"port" mapstructure:"port"`
	Timeout     time.Duration `yaml:"timeout" mapstructure:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout" mapstructure:"idle_timeout"`
}
