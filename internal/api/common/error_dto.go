package common

import "fmt"

type ApiError struct {
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func (apiError ApiError) Error() string {
	return fmt.Sprintf("[Code:%d]Error during execution request: %s", apiError.ErrorCode, apiError.ErrorMessage)
}
