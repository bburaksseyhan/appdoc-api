package utils

type ResponseResult struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// BadRequestError return ResponseError with bad_request status and messages
func Response(message string, data map[string]interface{}) *ResponseResult {
	return &ResponseResult{
		Message: message,
		Data:    data,
	}
}
