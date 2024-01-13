package http

import "net/http"

func ReturnErrorWithStatusCode(response http.ResponseWriter, statusCode int32, err error) {

}

func ReturnError(response http.ResponseWriter, err error) {

}
