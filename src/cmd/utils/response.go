package utils

type ResponseResult struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// Response and error_response.go file usage
func Response(message string, data map[string]interface{}) *ResponseResult {
	return &ResponseResult{
		Message: message,
		Data:    data,
	}
}
