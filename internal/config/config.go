package config

import "github.com/kelseyhightower/envconfig"

var AppConfig Config

type Config struct {
	Server struct {
		Port string `envconfig:"APP_PORT"`
	}
	Database struct {
		User     string `envconfig:"MONGO_USERNAME"`
		Password string `envconfig:"MONGO_PASSWORD"`
		Address  string `envconfig:"MONGO_URL"`
	}
	SearchEngine struct {
		Key string `envconfig:"MEILI_KEY"`
	}
}

func InitializeConfig() {
	err := envconfig.Process("", &AppConfig)
	if err != nil {
		panic(err)
	}
}
