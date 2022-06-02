package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Server struct {
		Port string `envconfig:"APP_PORT"`
	}
	Database struct {
		User     string
		Password string
		Address  string
	}
	SearchEngine struct {
		Address string
	}
}

func InitializeConfig() *Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
