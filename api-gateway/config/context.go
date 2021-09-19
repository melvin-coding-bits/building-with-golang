package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

type AppContextKey struct{}

type AppContext struct {
	Config *Config
	Logger *logrus.Logger
}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func InitContext() (*AppContext, error) {
	logger := initLogger()

	cnf, err := initConfig()
	if err != nil {
		logger.Error("error initializing config")
		return nil, err
	}
	logger.SetLevel(cnf.LogLevel)
	logger.Info("config initialized")

	return &AppContext{Config: cnf, Logger: logger}, nil
}
