package main

import (
	"os"
	"strconv"

	"github.com/bburakseyhann/appdoc-api/src/cmd/utils"
	"github.com/bburakseyhann/appdoc-api/src/pkg/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	config := read_configuration(read())

	// that func take config params.
	server.Initialize(config)
}

func read_configuration(config utils.Configuration) utils.Configuration {

	mongoUri := os.Getenv("MONGODB_URL")
	port := os.Getenv("SERVER_PORT")
	dbName := os.Getenv("DB_NAME")
	collection := os.Getenv("COLLECTION")
	appName := os.Getenv("APP_NAME")
	requestTimeOut, err := strconv.Atoi(os.Getenv("REQUEST_TIMEOUT"))

	if err != nil {
		logrus.Error("REQUEST_TIMEOUT can not be convert")
	}

	if mongoUri != "" || port != "" || dbName != "" || collection != "" || appName != "" {
		return utils.Configuration{
			App:      utils.Application{Name: appName, RequestTimeOut: requestTimeOut},
			Database: utils.DatabaseSetting{Url: mongoUri, DbName: dbName, Collection: collection},
			Server:   utils.ServerSettings{Port: port},
		}
	}

	// return config.yml variable
	return utils.Configuration{
		App:      utils.Application{Name: config.App.Name, RequestTimeOut: config.App.RequestTimeOut},
		Database: utils.DatabaseSetting{Url: config.Database.Url, DbName: config.Database.DbName, Collection: config.Database.Collection},
		Server:   utils.ServerSettings{Port: config.Server.Port},
	}
}

func read() utils.Configuration {
	//Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var config utils.Configuration

	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		logrus.Error("Unable to decode into struct, %v", err)
	}

	logrus.Warn("Config with variables %v", config)

	return config
}
