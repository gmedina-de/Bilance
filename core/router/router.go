package router

import (
	"net/http"
)

type Router interface {
	Init()
	Handle(writer http.ResponseWriter, request *http.Request)
}
