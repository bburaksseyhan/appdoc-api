package server

import (
	"fmt"

	"github.com/bburaksseyhan/appdoc-api/src/cmd/utils"
	"github.com/bburaksseyhan/appdoc-api/src/pkg/client/mongodb"
	"github.com/bburaksseyhan/appdoc-api/src/pkg/handler"
	"github.com/bburaksseyhan/appdoc-api/src/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Initialize(config utils.Configuration) {

	// Create a new instance of the logger. You can have any number of instances.
	var log = logrus.New()

	log.WithFields(logrus.Fields{
		"mongo_url":   config.Database.Url,
		"server_port": config.Server.Port,
		"db_name":     config.Database.DbName,
		"collection":  config.Database.Collection,
		"timeout":     config.App.Timeout,
	}).Info("\nConfiguration informations\n")

	logrus.Infof("Application Name %s is starting....", config.App.Name)

	client, err := mongodb.ConnectMongoDb(config.Database.Url)

	if err != nil {
		logrus.Fatal(err)
	}

	repository := repository.NewAppDocRepository(&config, client)
	handler := handler.NewAppDocHandler(client, repository, config)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	api := router.Group("api/v1")
	{
		api.GET("/health", handler.Healthcheck)

		api.POST("/appdoc/add", handler.Add)
		api.GET("/appdoc/list/:take", handler.List)
		api.GET("/appdoc/get/:id", handler.GetById)
		api.PUT("/appdoc/delete/:id", handler.Delete)
	}

	// PORT environment variable was defined.
	formattedUrl := fmt.Sprintf(":%s", config.Server.Port)

	router.Run(formattedUrl)
}
