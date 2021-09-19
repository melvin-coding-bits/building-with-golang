package config

import (
	"fmt"
	"os"

	"github.com/melvinodsa/build-with-golang/user-service/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//AppContext is the key access the app context from server context
type AppContextKey struct{}

//AppContext has the configuration, logger and database required by the server
type AppContext struct {
	//Config of the application
	Config *Config
	//Logger for logging
	Logger *logrus.Logger
	//DB to commuinicate with the database
	DB *gorm.DB
}

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func initDatabase(cnf Config, logger *logrus.Logger) (*gorm.DB, error) {
	/*
	 * Open the connection
	 * Do model migration
	 */
	//open the connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cnf.Database.Host, cnf.Database.User, cnf.Database.Password, cnf.Database.Name, cnf.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("error connecting to database")
		return nil, err
	}

	//do model migration
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Error("error doing automigration for user model")
		return nil, err
	}
	return db, nil
}

//InitAppContext initializes the app context
func InitContext() (*AppContext, error) {
	/*
	 * We will initialize the logger
	 * Then will initialize the config
	 * Will set the log level
	 * Then init the database
	 * Return the context instance
	 */
	//init the logger
	logger := initLogger()

	//init the config
	cnf, err := initConfig()
	if err != nil {
		logger.Error("error initializing config")
		return nil, err
	}

	//set the log level
	logger.SetLevel(cnf.LogLevel)
	logger.Info("config initialized")

	//init the database
	db, err := initDatabase(*cnf, logger)
	if err != nil {
		logger.Error("error initializing database")
		return nil, err
	}

	//return the context instance
	return &AppContext{Config: cnf, Logger: logger, DB: db}, nil
}
