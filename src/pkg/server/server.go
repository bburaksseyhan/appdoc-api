package server

import (
	"fmt"
	"net/http"

	"github.com/bburaksseyhan/appdoc-api/src/cmd/utils"
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
	}).Info("\nConfiguration informations\n")

	logrus.Infof("Application Name %s is starting....", config.App.Name)

	// Todo: Initialize MongoDb Client
	// Todo: Initialize Repositories and Handlers

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	// PORT environment variable was defined.
	formattedUrl := fmt.Sprintf(": %s", config.Server.Port)

	router.Run(formattedUrl)
}
