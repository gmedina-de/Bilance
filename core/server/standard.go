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
	log           log.Log
	authenticator authenticator.Authenticator
	router        router.Router
}

func Standard(log log.Log, authenticator authenticator.Authenticator, router router.Router) Server {
	return &standard{log, authenticator, router}
}

func (r *standard) Serve() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", r.ServeHTTP)
	r.log.Info("Starting server http://localhost:%d", config.ServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), nil)
	if err != nil {
		r.log.Critical(err.Error())
	}
}

func (r *standard) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	w := &statusWriter{ResponseWriter: writer, status: 200}
	if r.authenticator == nil || r.authenticator.Authenticate(w, request) {
		r.router.Handle(w, request)
	}
	r.log.Debug("%s %s -> %d", request.Method, request.URL, w.status)
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
