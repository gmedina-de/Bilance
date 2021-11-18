package service

import (
	"fmt"
	"net/http"
	"strconv"
)

type Server interface {
	Start()
	Stop()
}

const (
	Port Setting = iota
)

type defaultServer struct {
	router Router
	log    Log
	port   int
}

func HttpServer(router Router, log Log, configuration Configuration) Server {
	var port int
	var error error
	port, error = strconv.Atoi(configuration.Get(Port))
	if error != nil {
		port = 8080
	}
	return &defaultServer{router: router, log: log, port: port}
}

func (this *defaultServer) Start() {
	http.HandleFunc("/", this.router.ServeHTTP)
	this.log.Info("Listening to http://localhost:%d", this.port)
	http.ListenAndServe(fmt.Sprintf(":%d", this.port), nil)
}

func (this *defaultServer) Stop() {
}
