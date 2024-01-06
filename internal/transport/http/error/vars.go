package error

import "net/http"

var (
	ErrXRealIPRequired = NewApiError(
		ErrXRealIPRequiredCode,
		http.StatusTooManyRequests,
		"X-Real-Ip required",
		"X-Real-Ip header is required",
	)

	ErrTooManyRequests = NewApiError(
		http.StatusTooManyRequests,
		http.StatusTooManyRequests,
		"Too many requests",
		"",
	)

	ErrInternal = NewApiError(
		http.StatusInternalServerError,
		http.StatusInternalServerError,
		"Internal server error",
		"",
	)
)
