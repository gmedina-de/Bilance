package router

import (
	"net/http"
)

type Router interface {
	Handle(writer http.ResponseWriter, request *http.Request)
}
