package server

import (
	"net/http"
)

type Server interface {
	SetBasePath(basePath string)
	Get(route string, handler http.HandlerFunc)
	Post(route string, handler http.HandlerFunc)
	Start()
}
