package server

import (
	"fmt"
	"genuine/config"
	"genuine/core/authenticator"
	"genuine/core/log"
	"genuine/core/router"
	"net/http"
)

type standard struct {
	Log           log.Log
	Authenticator authenticator.Authenticator
	Router        router.Router
}

func Standard() Server {
	return &standard{}
}

func (r *standard) Start() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", r.ServeHTTP)
	r.Log.Info("Starting server http://localhost:%d", config.ServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), nil)
	if err != nil {
		r.Log.Critical(err.Error())
	}
}

func (r *standard) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	w := &statusWriter{ResponseWriter: writer, status: 200}
	if r.Authenticator == nil || r.Authenticator.Authenticate(w, request) {
		r.Router.Handle(w, request)
	}
	r.Log.Debug("%s %s -> %d", request.Method, request.URL, w.status)
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
