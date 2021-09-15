package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	Server   Server
	LogLevel logrus.Level
	Database Database
	Verison  Verison
}

func (s Server) GetPort() string {
	return fmt.Sprintf(":%s", s.Port)
}

func initConfig() (*Config, error) {
	return &Config{
		Server: Server{
			Port: "3000",
		},
		LogLevel: logrus.DebugLevel,
		Database: Database{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			Name:     "building_with_golang",
		},
		Verison: V1,
	}, nil
}
