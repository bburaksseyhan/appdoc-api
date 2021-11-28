# appdoc-api

In this branch you'll see;

* Project layout :building_construction:
* Implemented <b>GIN Web Framework</b> :genie_man:
* Read configuration file using with <b>Viper</b> :snake: and OS package :package:
* You'll learn <b>logrus</b> implementation
* Implement the MongoDB using with <b>mongo</b> and <b>options</b> packages :package:
* Working with Repository Pattern and  Factory Design Pattern (includes repository and handler)
* Using OS packages :package:
* Writing custom response.go :package:
* Docker file and docker compose file added :ferry:

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

### response.go
```
type ResponseResult struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// Response return some of information
func Response(message string, data map[string]interface{}) *ResponseResult {
	return &ResponseResult{
		Message: message,
		Data:    data,
	}
}

<b>usage</b> 

c.IndentedJSON(http.StatusOK, utils.Response("Pong", map[string]interface{}{"Data": "The MongoDB client is working successfully", "Date": time.Local}))
```

### docker file
```
	FROM golang:1.16-alpine as build-env
	WORKDIR /app
	COPY go.mod go.sum ./
	RUN go mod download
	COPY . ./
	RUN  go build -o /appdoc-api github.com/bburaksseyhan/appdoc-api/src/cmd/api   

	FROM alpine:3.14

	RUN apk update \
		&& apk upgrade\
		&& apk add --no-cache tzdata curl

	#RUN apk --no-cache add bash
	ENV TZ Europe/Istanbul

	WORKDIR /app
	COPY --from=build-env /appdoc-api .
	COPY --from=build-env /app/src/cmd/api /app/

	EXPOSE 80
	CMD [ "./appdoc-api" ]
```

### docker compose file
```
version: "3.8"
  
services:
    mongodb:
      image : mongo
      container_name: mongodb
      ports:
      - 27017:27017
      healthcheck:
        test:
        - CMD
        - mongo
        - --eval
        - "db.adminCommand('ping')"
      restart: unless-stopped
    
    appdocapi:
      build:
        context: .
        dockerfile: ./dockerfile
      ports: 
        - 8080:8080
      restart: on-failure
      env_file:
        - .env
      depends_on:
        mongodb:
          condition: service_healthy
```

### docker commands

```
 docker compose up
```

### mongodb cli and queries

```
 docker ps -a
 docker exec -it container_id mongo
 show dbs
 use AppDb
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
