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

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(app.config.App.Timeout)*time.Second)
	defer ctxErr()

	if ctxErr != nil {
		logrus.Error("somethig wrong!!!", ctxErr)
	}

	if err := app.client.Ping(ctx, nil); err != nil {
		utils.InternalServerError("Status unhealth", err, map[string]interface{}{"Data": "Please check the Client", "Time": time.Local})
	}

	c.IndentedJSON(http.StatusOK, utils.Response("Pong", map[string]interface{}{"Data": "The MongoDB client is working successfully", "Date": time.Local}))
}

func (app *appDocHandler) Add(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(app.config.App.Timeout)*time.Second)
	defer ctxErr()

	var appEntity *entity.AppDoc

	appEntity = &entity.AppDoc{
		CompanyName:  c.Param("company_name"),
		AppName:      c.Param("app_name"),
		AppVersion:   c.Param("app_version"),
		Domain:       c.Param("domain"),
		EmailAddress: c.Param("email_address"),
		IpAddress:    c.Param("ip_address"),
		Url:          c.Param("url"),
		Country:      c.Param("country"),
	}

	entity := entity.AppDoc(*appEntity)

	oId, err := app.appDocRepository.Add(entity, ctx)
	if err != nil {
		utils.BadRequestError("AppDoc_Handler_Add", err, map[string]interface{}{"Data": entity})
	}

	c.IndentedJSON(http.StatusCreated, utils.Response("AppDoc_Handler_Add", map[string]interface{}{"OId": oId}))
}

func (app *appDocHandler) List(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(app.config.App.Timeout)*time.Second)
	defer ctxErr()

	var appDocsModel []*model.AppDoc

	tParam := c.Param("take")
	take, err := strconv.Atoi(tParam)
	if err != nil {
		logrus.Fatal("Take parameter can not be converted.")
	}

	logrus.Infof("Take %d", take)

	result, err := app.appDocRepository.List(take, ctx)
	if err != mongo.ErrNilCursor {
		utils.BadRequestError("AppDoc_Handler_List", err, map[string]interface{}{"Data": take})
	}
	logrus.Infof("Len %d", len(result))

	//convert to entity to model
	for _, item := range result {
		appDocsModel = append(appDocsModel, (*model.AppDoc)(item))
	}

	c.IndentedJSON(http.StatusOK, utils.Response("AppDoc_Handler_List", map[string]interface{}{"Data": appDocsModel}))
}

func (app *appDocHandler) GetById(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(app.config.App.Timeout)*time.Second)
	defer ctxErr()

	id := c.Param("id")

	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.BadRequestError("AppDoc_Handler_GetById ObjectId couldn't convert it.", err, map[string]interface{}{"Data": id})
	}

	logrus.Infof("OId %s", oId)

	result, err := app.appDocRepository.GetById(oId, ctx)
	if err != mongo.ErrNilCursor {
		utils.BadRequestError("AppDoc_Handler_GetById", err, map[string]interface{}{"Data": id})
	}

	if result == nil {
		utils.NotFoundRequestError("AppDoc_Handler_GetById", err, map[string]interface{}{"Data": id})
	}

	c.IndentedJSON(http.StatusOK, utils.Response("AppDoc_Handler_GetById", map[string]interface{}{"Data": result}))
}

func (app *appDocHandler) Delete(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(app.config.App.Timeout)*time.Second)
	defer ctxErr()

	id := c.Param("id")

	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.BadRequestError("AppDoc_Handler_GetById ObjectId couldn't convert it.", err, map[string]interface{}{"Data": id})
	}

	logrus.Infof("OId %s", oId)

	result, err := app.appDocRepository.Delete(oId, ctx)
	if err != nil {
		utils.BadRequestError("AppDoc_Handler_Delete", err, map[string]interface{}{"Data": id})
	}

	c.IndentedJSON(http.StatusOK, utils.Response("AppDoc_Handler_Delete", map[string]interface{}{"Data": result}))
}
