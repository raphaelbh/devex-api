package config

import (
	"log/slog"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT" default:"8080"`
	}
	Database struct {
		Url string `envconfig:"DATABASE_URL" default:"host=localhost user=brekke password=brekke dbname=brekke port=5432"`
	}
	Authentication struct {
		Key string `envconfig:"AUTHENTICATION_KEY" default:"auth-key"`
	}
	Encryption struct {
		Key string `envconfig:"SECURITY_ENCRYPTION_KEY" default:"f1a3c47d9e20b8653fa472d9bf38ab12"`
	}
}

var (
	conf   config
	logger = slog.Default()
)

func Default() config {
	return conf
}

func init() {
	var c config
	err := envconfig.Process("", &c)
	if err != nil {
		logger.Error("Unable to get environment variables", slog.String("error", err.Error()))
		panic("Unable to get environment variables: " + err.Error())
	}

	conf = c
	logger.Info("Environment variables loaded")
}
