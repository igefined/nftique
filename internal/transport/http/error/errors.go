package error

type ApiError struct {
	// Code mean internal code
	Code       int    `json:"code"`
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func NewApiError(code int, statusCode int, message string) *ApiError {
	return &ApiError{Code: code, StatusCode: statusCode, Message: message}
}

func (e *ApiError) Error() string {
	return e.Message
}
