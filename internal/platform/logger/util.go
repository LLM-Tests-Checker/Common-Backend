package logger

import (
	"net/http"
)

func GetLogger(request *http.Request) Logger {
	logger := request.Context().Value(VariableLogger)
	return logger.(Logger)
}
