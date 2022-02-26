package controllers

import (
	"genuine/core/http"
)

type Controller interface {
	Routes() map[string]http.Handler
}
