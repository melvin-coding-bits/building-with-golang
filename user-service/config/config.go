//Package config has the utilities required for defining and
//initializing the server configuration
package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

//Server has the config required by the api server
type Server struct {
	//Port on which the api server should run
	Port string
}

//Database has the config required by the database
type Database struct {
	//Host of the database
	Host string
	//Port on which the database is listening
	Port string
	//User to access the database
	User string
	//Password to access the database
	Password string
	//Name of the database to access
	Name string
}

//Config has the config required by the api server
type Config struct {
	//Server config
	Server Server
	//LogLevel of the server logs
	LogLevel logrus.Level
	//Database config
	Database Database
	//Version of the api
	Verison Verison
}

//GetPort returns the port with required notation for listening ie. :port
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
