package utils

import "net/http"

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
