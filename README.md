# appdoc-api

In this branch you'll see;

* Project layout :building_construction:
* Implemented <b>GIN Web Framework</b> :genie_man:
* Read configuration file using with <b>Viper</b> :snake: and OS package :package:
* You'll learn <b>logrus</b> implementation
* Implement the MongoDB using with <b>mongo</b> and <b>options</b> packages :package:
* Working with Repository Pattern and  Factory Design Pattern (includes repository and handler)

required packages :package:
```
    go get -u github.com/sirupsen/logrus v1.8.1
	go get -u github.com/spf13/viper v1.9.0
    go get -u github.com/gin-gonic/gin
    go get -u go.mongodb.org/mongo-driver/mongo
    go get -u go.mongodb.org/mongo-driver/mongo/options
    go get -u go.mongodb.org/mongo-driver/bson/primitive
    go get -u go.mongodb.org/mongo-driver/bson
```

```
 go mod init github.com/bburakseyhann/appdoc-api
```

```
 go run main.go
```


```
 curl http://localhost:8080/health
```
