package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Port string
}

type Config struct {
	Server   Server
	LogLevel logrus.Level
	Verison  Verison
}

func (s Server) GetPort() string {
	return fmt.Sprintf(":%s", s.Port)
}

func initConfig() (*Config, error) {
	return &Config{
		Server: Server{
			Port: "8080",
		},
		LogLevel: logrus.DebugLevel,
		Verison:  V1,
	}, nil
}
