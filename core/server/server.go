package server

import (
	"net/http"
)

type Server interface {
	Get(route string, handler http.HandlerFunc)
	Post(route string, handler http.HandlerFunc)
	Start()
}
