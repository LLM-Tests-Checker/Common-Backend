package common

type ApiError struct {
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
