package config

import (
	"flag"

	"github.com/caarlos0/env/v10"
	"go.uber.org/zap"
)

// InitLogger configures zap logger
func InitLogger(debug bool, projectID string) (*zap.Logger, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.LevelKey = "severity"
	zapConfig.EncoderConfig.MessageKey = "message"

	if debug {
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, err := zapConfig.Build(zap.Fields(
		zap.String("projectID", projectID),
	))

	if err != nil {
		return nil, err
	}

	return logger, nil
}

type Config struct {
	PathApp   string `env:"PATHAPP"`
	timeExec  int    `env:"timeexec"`
	PathCheck string `env:"pathcheck"`

	AppName string `env:"APP_NAME" envDefault:"AppCheckPathExist"`
	Debug   bool   `env:"DEBUG"`
	Logger  *zap.Logger
}

// InitConfig initialises config, first from flags, then from env, so that env overwrites flags
func InitConfig() (*Config, error) {
	var cfg Config

	flag.StringVar(&cfg.PathApp, "path_app", "", "path_app=1.exe")
	flag.StringVar(&cfg.PathCheck, "path_check", "", "path_check=c:\\123")
	flag.Parse()

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
