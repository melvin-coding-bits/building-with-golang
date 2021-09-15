package config

import (
	"fmt"
	"os"

	"github.com/melvinodsa/build-with-golang/user-service/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppContextKey struct{}

type AppContext struct {
	Config *Config
	Logger *logrus.Logger
	DB     *gorm.DB
}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func initDatabase(cnf Config, logger *logrus.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cnf.Database.Host, cnf.Database.User, cnf.Database.Password, cnf.Database.Name, cnf.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("error connecting to database")
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Error("error doing automigration for user model")
		return nil, err
	}
	return db, nil
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

	db, err := initDatabase(*cnf, logger)
	if err != nil {
		logger.Error("error initializing database")
		return nil, err
	}

	return &AppContext{Config: cnf, Logger: logger, DB: db}, nil
}
