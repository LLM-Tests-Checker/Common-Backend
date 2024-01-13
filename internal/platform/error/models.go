package error

import "fmt"

type ErrorCode = int32

type BackendError struct {
	Code    ErrorCode
	Message string

	StatusCode int
}

func NewBackendError(
	code ErrorCode,
	message string,
	statusCode int,
) *BackendError {
	return &BackendError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

func (err *BackendError) Error() string {
	return fmt.Sprintf("[BackendError] Code: %d, Message: %s", err.Code, err.Message)
}
