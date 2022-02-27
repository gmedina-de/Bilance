package controllers

import "net/http"

type Request struct {
	*http.Request
	http.ResponseWriter
}

type Response = map[string]any

type Handler func(r Request) Response
