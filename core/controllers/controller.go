package controllers

import "net/http"

type Controller interface {
	Before(request *http.Request, writer http.ResponseWriter, name string)
	Routes() map[string]string
	After()
}
