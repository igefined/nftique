package error

type ApiError struct {
	// Code mean internal code
	Code        int    `json:"code"`
	StatusCode  int    `json:"-"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func NewApiError(code int, statusCode int, message, description string) *ApiError {
	return &ApiError{Code: code, StatusCode: statusCode, Message: message, Description: description}
}

func (e *ApiError) Error() string {
	return e.Message
}
