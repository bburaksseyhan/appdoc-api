package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppDocHandler interface {
	Healthcheck(*gin.Context)
}

type appDocHandler struct {
	client *mongo.Client
}

func NewAppDocHandler(client *mongo.Client) AppDocHandler {
	return &appDocHandler{client: client}
}

func (app *appDocHandler) Healthcheck(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer ctxErr()

	if ctxErr != nil {
		logrus.Error("somethig wrong!!!", ctxErr)
	}

	if err := app.client.Ping(ctx, nil); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "unhealty"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "pong"})
}
