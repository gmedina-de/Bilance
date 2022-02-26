package http

import "net/http"

type Request = *http.Request

type Response = map[string]any

type Handler = func(r Request) Response
