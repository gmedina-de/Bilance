package server

import (
	"fmt"
	"genuine/core/authenticator"
	"genuine/core/log"
	"genuine/core/router"
	"net/http"
)

type standard struct {
	Log           log.Log
	Authenticator authenticator.Authenticator
	Router        router.Router
	port          int
}

func Standard() Server {
	return &standard{port: 8080}
}

func (r *standard) Start() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", r.ServeHTTP)
	r.Log.Info("Listening to http://localhost:%d", r.port)
	http.ListenAndServe(fmt.Sprintf(":%d", r.port), nil)
}

func (r *standard) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.Log.Debug("%s %s", request.Method, request.URL)
	if r.Authenticator.Authenticate(writer, request) {
		r.Router.Handle(writer, request)
	}
}
