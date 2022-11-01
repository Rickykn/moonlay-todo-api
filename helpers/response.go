package helpers

type JsonResponse struct {
	Code    int
	Message string
	Data    interface{}
	Error   bool
}

func HandlerSuccess(statusCode int, message string, data interface{}) *JsonResponse {
	return &JsonResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
		Error:   false,
	}
}

func HandlerError(statusCode int, message string, data interface{}) *JsonResponse {
	return &JsonResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
		Error:   true,
	}
}
