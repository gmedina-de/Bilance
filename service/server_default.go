package service

import (
	"fmt"
	"net/http"
)

type defaultServer struct {
	router Router
	log    Log
	port   int
}

func DefaultServer(router Router, log Log) Server {
	return &defaultServer{router: router, log: log, port: 8080}
}

func (this *defaultServer) Start() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", this.router.ServeHTTP)
	this.log.Info("Listening to http://localhost:%d", this.port)
	http.ListenAndServe(fmt.Sprintf(":%d", this.port), nil)
}

func (this *defaultServer) Stop() {
}
