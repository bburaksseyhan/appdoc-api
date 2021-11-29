package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bburaksseyhan/appdoc-api/src/cmd/utils"
	"github.com/bburaksseyhan/appdoc-api/src/pkg/entity"
	"github.com/bburaksseyhan/appdoc-api/src/pkg/model"
	"github.com/bburaksseyhan/appdoc-api/src/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppDocHandler interface {
	Healthcheck(*gin.Context)

	Add(*gin.Context)
	List(*gin.Context)
	GetById(*gin.Context)
	Delete(*gin.Context)
}

type appDocHandler struct {
	client           *mongo.Client
	appDocRepository repository.AppDocRepository
	config           utils.Configuration
}

func NewAppDocHandler(client *mongo.Client, repo repository.AppDocRepository, config utils.Configuration) AppDocHandler {
	return &appDocHandler{client: client, appDocRepository: repo, config: config}
}

func (app *appDocHandler) Healthcheck(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(app.config.App.RequestTimeOut)*time.Second)
	defer ctxErr()

	if ctxErr != nil {
		logrus.Error("somethig wrong!!!", ctxErr)
	}

	if err := app.client.Ping(ctx, nil); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "unhealty"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "pong"})
}

func (app *appDocHandler) Add(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(app.config.App.RequestTimeOut)*time.Second)
	defer ctxErr()

	var appModel *model.AppDoc

	//get parameter
	company_name := c.Param("company_name")
	app_name := c.Param("app_name")
	app_version := c.Param("app_version")
	domain := c.Param("domain")
	email_address := c.Param("email_address")
	ip_address := c.Param("ip_address")
	url := c.Param("url")
	country := c.Param("country")

	appModel = &model.AppDoc{
		Id:           [12]byte{},
		CompanyName:  company_name,
		AppName:      app_name,
		AppVersion:   app_version,
		Domain:       domain,
		EmailAddress: email_address,
		IpAddress:    ip_address,
		Url:          url,
		Country:      country,
	}

	entity := entity.AppDoc(*appModel)

	oId, err := app.appDocRepository.Add(entity, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"add result": err.Error})
	}

	c.JSON(http.StatusOK, gin.H{"add result": oId})
}

func (app *appDocHandler) List(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(app.config.App.RequestTimeOut)*time.Second)
	defer ctxErr()

	var appDocsModel []*model.AppDoc

	tParam := c.Param("take")
	take, _ := strconv.Atoi(tParam)

	result, err := app.appDocRepository.List(take, ctx)
	if err != mongo.ErrNilCursor {
		c.JSON(http.StatusBadRequest, gin.H{"list result": err.Error()})
	}

	//convert to entity to model
	for _, item := range result {
		appDocsModel = append(appDocsModel, (*model.AppDoc)(item))
	}

	c.JSON(http.StatusOK, gin.H{"list result": appDocsModel})
}

func (app *appDocHandler) GetById(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(app.config.App.RequestTimeOut)*time.Second)
	defer ctxErr()

	id := c.Param("id")

	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Error("can not convert to id")
		c.JSON(http.StatusBadRequest, gin.H{"GetById": err.Error()})
	}

	result, err := app.appDocRepository.GetById(oId, ctx)
	if err != mongo.ErrNilCursor {
		c.JSON(http.StatusBadRequest, gin.H{"GetById result": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"GetById result": result})
}

func (app *appDocHandler) Delete(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(app.config.App.RequestTimeOut)*time.Second)
	defer ctxErr()

	id := c.Param("id")

	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logrus.Error("can not convert to id")
		c.JSON(http.StatusBadRequest, gin.H{"Delete": err.Error()})
	}

	result, err := app.appDocRepository.Delete(oId, ctx)
	if err != mongo.ErrNilCursor {
		c.JSON(http.StatusBadRequest, gin.H{"Delete result": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"Delete result": result})
}
