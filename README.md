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

### error_response.go
```
type ResponseError struct {
	Message string                 `json:"message"`
	Status  int                    `json:"status"`
	Error   string                 `json:"error"`
	Data    map[string]interface{} `json:"data"`
}

// BadRequestError return ResponseError with bad_request status and messages
func BadRequestError(message string, err error, data map[string]interface{}) *ResponseError {
	return &ResponseError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
		Data:    data,
	}
}

// NotFoundRequestError return ResponseError with not_found status and messages
func NotFoundRequestError(message string, err error, data map[string]interface{}) *ResponseError {
	return &ResponseError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
		Data:    data,
	}
}

//mcustom error return ResponseError with internal_server status and messages
func InternalServerError(message string, err error, data map[string]interface{}) *ResponseError {
	return &ResponseError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server",
		Data:    data,
	}
}

<b>usage</b> 
utils.InternalServerError("Status unhealth", err, map[string]interface{}{"Data": "Please check the Client", "Time": time.Local})

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
