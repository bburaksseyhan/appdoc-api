package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Initialize() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Application is starting....")
	// Todo: Configure Application Settings with Viper
	// Todo: Initialize MongoDb Client
	// Todo: Initialize Repositories and Handlers

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	// PORT environment variable was defined.
	router.Run(":8080")
}
