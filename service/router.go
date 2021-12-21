package service

import (
	"net/http"
)

type Router interface {
	Get(route string, handler Handler)
	Post(route string, handler Handler)
	Start()
	http.Handler
}

type Handler func(http.ResponseWriter, *http.Request)
