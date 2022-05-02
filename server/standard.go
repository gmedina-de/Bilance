package server

import (
	"flag"
	"fmt"
	"genuine/log"
	"genuine/router"
	"net/http"
)

var port = flag.Int("port", 8080, "application port")

type standard struct {
	log    log.Log
	router router.Router
}

func Standard(log log.Log, router router.Router) Server {
	return &standard{log, router}
}

func (r *standard) Serve() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", r.ServeHTTP)
	r.log.Info("Starting server http://localhost:%d", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		r.log.Critical(err.Error())
	}
}

func (r *standard) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	w := &statusWriter{ResponseWriter: writer, status: 200}
	r.router.Handle(w, request)
	r.log.Debug("%s %s -> %d", request.Method, request.URL, w.status)
}

// in order to know code status
type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
