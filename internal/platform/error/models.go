package error

import "fmt"

type ErrorCode = int32

type BackendError struct {
	Code    ErrorCode
	Message string

	StatusCode int32
}

func NewBackendError(
	code ErrorCode,
	message string,
	statusCode int32,
) *BackendError {
	return &BackendError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

func (err BackendError) Error() string {
	return fmt.Sprintf("[BackendError] Code: %d, Message: %s", err.Code, err.Message)
}

func Wrap(
	err error,
	code ErrorCode,
	message string,
	statusCode int32,
) *BackendError {
	return &BackendError{
		Code:       code,
		Message:    fmt.Sprintf("%s: %s", message, err.Error()),
		StatusCode: statusCode,
	}
}
