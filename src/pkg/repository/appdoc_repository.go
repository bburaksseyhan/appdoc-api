package repository

import (
	"context"
	"log"

	"github.com/bburakseyhann/appdoc-api/src/cmd/utils"
	"github.com/bburakseyhann/appdoc-api/src/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppDocRepository interface {
	Add(appDoc entity.AppDoc, ctx context.Context) (primitive.ObjectID, error)
	List(count int, ctx context.Context) ([]*entity.AppDoc, error)
	GetById(oId primitive.ObjectID, ctx context.Context) (*entity.AppDoc, error)
	Delete(oId primitive.ObjectID, ctx context.Context) (int64, error)
}

type appDocRepository struct {
	client *mongo.Client
	config *utils.Configuration
}

func NewAppDocRepository(config *utils.Configuration, client *mongo.Client) AppDocRepository {
	return &appDocRepository{config: config, client: client}
}

func (app *appDocRepository) Add(appDoc entity.AppDoc, ctx context.Context) (primitive.ObjectID, error) {

	collection := app.client.Database(app.config.Database.DbName).Collection(app.config.Database.Collection)

	insertResult, err := collection.InsertOne(ctx, appDoc)

	if err != mongo.ErrNilCursor {
		return primitive.NilObjectID, err
	}

	if oidResult, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		return oidResult, nil
	} else {
		return primitive.NilObjectID, err
	}
}

func (app *appDocRepository) List(count int, ctx context.Context) ([]*entity.AppDoc, error) {

	findOptions := options.Find()
	findOptions.SetLimit(int64(count))

	collection := app.client.Database(app.config.Database.DbName).Collection(app.config.Database.Collection)

	cursor, err := collection.Find(ctx, findOptions)
	if err != nil {
		return nil, err
	}

	var appDocs []*entity.AppDoc
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem entity.AppDoc
		if err := cursor.Decode(&elem); err != nil {
			log.Fatal(err)
			return nil, err
		}

		appDocs = append(appDocs, &elem)
	}

	cursor.Close(ctx)

	return appDocs, mongo.ErrNilCursor
}

func (app *appDocRepository) GetById(oId primitive.ObjectID, ctx context.Context) (*entity.AppDoc, error) {

	collection := app.client.Database(app.config.Database.DbName).Collection(app.config.Database.Collection)

	filter := bson.D{primitive.E{Key: "_id", Value: oId}}

	var appDoc *entity.AppDoc

	collection.FindOne(ctx, filter).Decode(&appDoc)

	return appDoc, bson.ErrDecodeToNil
}

func (app *appDocRepository) Delete(oId primitive.ObjectID, ctx context.Context) (int64, error) {

	collection := app.client.Database(app.config.Database.DbName).Collection(app.config.Database.Collection)
	filter := bson.D{primitive.E{Key: "_id", Value: oId}}

	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return 0, bson.ErrDecodeToNil
	}

	return result.DeletedCount, nil
}
