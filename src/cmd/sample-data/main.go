package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bburaksseyhan/appdoc-api/src/pkg/client/mongodb"
	"github.com/bburaksseyhan/appdoc-api/src/pkg/entity"

	"github.com/sirupsen/logrus"
)

func main() {

	appDocJSON, err := os.Open("data.json")

	if err != nil {
		logrus.Fatal("data.json an error occurred", err)
	}

	defer appDocJSON.Close()

	appDocs := []entity.AppDoc{}

	byteValue, _ := ioutil.ReadAll(appDocJSON)

	//unmarshall data
	if err := json.Unmarshal(byteValue, &appDocs); err != nil {
		logrus.Error("unmarshall an error occurred", err)
	}

	logrus.Info("Data\n", len(appDocs))

	//import mongo client
	client, err := mongodb.ConnectMongoDb("mongodb://localhost:27017")
	if err != nil {
		logrus.Fatal(err)
	}

	defer client.Disconnect(context.TODO())

	collection := client.Database("AppDb").Collection("applications")
	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		logrus.Fatal(err.Error())
	}

	logrus.Info("MongoDb Client connection success")

	logrus.Warn("Total data count:", &appDocs)

	for _, item := range appDocs {
		collection.InsertOne(context.TODO(), item)
	}

	logrus.Info("Data import finished...")
}
