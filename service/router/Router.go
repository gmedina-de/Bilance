package router

import (
	"net/http"
)

type Router interface {
	Get(route string, handler Handler)
	Post(route string, handler Handler)
	http.Handler
}
