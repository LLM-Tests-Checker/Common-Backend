package common

import "fmt"

type ApiError struct {
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`

	httpStatusCode *int32
}

func NewApiError(errorCode int32, message string, defaultStatusCode *int32) ApiError {
	return ApiError{
		ErrorCode:      errorCode,
		ErrorMessage:   message,
		httpStatusCode: defaultStatusCode,
	}
}

func (apiError ApiError) Error() string {
	return fmt.Sprintf("[Code:%d]Error during execution request: %s", apiError.ErrorCode, apiError.ErrorMessage)
}

func (apiError ApiError) GetDefaultStatusCode() *int32 {
	return apiError.httpStatusCode
}
