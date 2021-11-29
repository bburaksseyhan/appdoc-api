# appdoc-api

<a href="https://codeclimate.com/github/bburaksseyhan/appdoc-api/maintainability"><img src="https://api.codeclimate.com/v1/badges/b93f81c081921cb109df/maintainability" /></a>
[![license](https://img.shields.io/github/license/bburaksseyhan/appdoc-api.svg)](LICENSE)
[![Go language](https://img.shields.io/badge/language-Go-blue.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bburaksseyhan/appdoc-api)](https://goreportcard.com/report/github.com/bburaksseyhan/appdoc-api)

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
    
    appdoc-api:
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
```

```
 show dbs
```

```
 use AppDb
```

```
 show collections
```

```
 db.applications.find({}).pretty()
```

```
 db.applications.find({"_id":ObjectId("61a476e6039fa0ff1792b7ff")}).pretty()
```


### run main.go

```
 go mod init github.com/bburakseyhann/appdoc-api
```

```
 go run main.go
```

```

 curl http://localhost:8080/api/v1/health

 {
    "message": "Pong",
    "data": {
        "Data": "The MongoDB client is working successfully",
        "Date": {}
    }
}

 curl http://localhost:8080/api/v1/appdoc/list/1

 {
    "message": "AppDoc_Handler_List",
    "data": {
        "Data": [
            {
                "company_name": "Rhyzio",
                "app_name": "Opela",
                "app_version": "9.9",
                "domain": "cisco.com",
                "email_address": "slinny0@ycombinator.com",
                "ip_address": "123.178.5.111",
                "url": "http://istockphoto.com/vitae/nisi.jpg?vulputate=fusce&justo=congue&in=diam&blandit=id&ultrices=ornare&enim=imperdiet&lorem=sapien&ipsum=urna&dolor=pretium&sit=nisl&amet=ut&consectetuer=volutpat&adipiscing=sapien&elit=arcu&proin=sed&interdum=augue&mauris=aliquam&non=erat&ligula=volutpat&pellentesque=in&ultrices=congue&phasellus=etiam&id=justo&sapien=etiam&in=pretium&sapien=iaculis&iaculis=justo&congue=in&vivamus=hac&metus=habitasse&arcu=platea&adipiscing=dictumst&molestie=etiam&hendrerit=faucibus&at=cursus&vulputate=urna&vitae=ut&nisl=tellus&aenean=nulla&lectus=ut&pellentesque=erat&eget=id&nunc=mauris&donec=vulputate&quis=elementum&orci=nullam&eget=varius&orci=nulla&vehicula=facilisi&condimentum=cras&curabitur=non&in=velit&libero=nec&ut=nisi&massa=vulputate&volutpat=nonummy&convallis=maecenas&morbi=tincidunt&odio=lacus&odio=at&elementum=velit&eu=vivamus&interdum=vel&eu=nulla&tincidunt=eget&in=eros&leo=elementum&maecenas=pellentesque&pulvinar=quisque&lobortis=porta&est=volutpat&phasellus=erat&sit=quisque&amet=erat&erat=eros&nulla=viverra&tempus=eget&vivamus=congue&in=eget&felis=semper&eu=rutrum&sapien=nulla&cursus=nunc&vestibulum=purus&proin=phasellus&eu=in&mi=felis&nulla=donec&ac=semper&enim=sapien&in=a&tempor=libero",
                "country": "France"
            }
        ]
    }
}

 curl http://localhost:8080/api/v1/appdoc/get/61a4ac47561122a1f7537dfd


{
    "message": "AppDoc_Handler_GetById",
    "data": {
        "Data": {
            "company_name": "Abatz",
            "app_name": "Bamity",
            "app_version": "0.92",
            "domain": "boston.com",
            "email_address": "ehavesidesj@sfgate.com",
            "ip_address": "231.91.24.100",
            "url": "http://prnewswire.com/ut/at/dolor/quis.aspx?aliquam=ut&lacus=dolor&morbi=morbi&quis=vel&tortor=lectus&id=in&nulla=quam&ultrices=fringilla&aliquet=rhoncus&maecenas=mauris&leo=enim&odio=leo&condimentum=rhoncus&id=sed&luctus=vestibulum&nec=sit&molestie=amet&sed=cursus&justo=id&pellentesque=turpis&viverra=integer&pede=aliquet&ac=massa&diam=id&cras=lobortis&pellentesque=convallis&volutpat=tortor&dui=risus&maecenas=dapibus&tristique=augue&est=vel&et=accumsan&tempus=tellus&semper=nisi&est=eu&quam=orci&pharetra=mauris&magna=lacinia&ac=sapien&consequat=quis&metus=libero&sapien=nullam&ut=sit&nunc=amet&vestibulum=turpis&ante=elementum&ipsum=ligula&primis=vehicula&in=consequat&faucibus=morbi&orci=a&luctus=ipsum&et=integer&ultrices=a&posuere=nibh&cubilia=in&curae=quis&mauris=justo&viverra=maecenas&diam=rhoncus&vitae=aliquam&quam=lacus&suspendisse=morbi&potenti=quis&nullam=tortor&porttitor=id&lacus=nulla&at=ultrices&turpis=aliquet&donec=maecenas&posuere=leo",
            "country": "France"
        }
    }
}

 curl http://localhost:8080/api/v1/appdoc/delete/61a4ac47561122a1f7537dfc

{
    "message": "AppDoc_Handler_Delete",
    "data": {
        "Data": 1
    }
}
```
