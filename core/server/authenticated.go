package server

import (
	"fmt"
	"homecloud/core/authenticator"
	"homecloud/core/log"
	"net/http"
	"strings"
)

type authenticated struct {
	log           log.Log
	authenticator authenticator.Authenticator
	routes        map[string]http.HandlerFunc
	port          int
	basePath      string
}

func Authenticated(log log.Log, authenticator authenticator.Authenticator) Server {
	return &authenticated{log: log, authenticator: authenticator, routes: make(map[string]http.HandlerFunc), port: 8080}
}

func (r *authenticated) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.log.Debug("%s %s", request.Method, request.URL)
	if r.authenticator.Authenticate(writer, request) {
		handler, found := r.routes[strings.ToUpper(request.Method)+" "+request.URL.Path]
		if found {
			handler(writer, request)
		} else {
			fmt.Fprintf(writer, "No route for %s found!", request.URL.Path)
		}
	}
}

func (r *authenticated) SetBasePath(basePath string) {
	r.basePath = basePath
}

func (r *authenticated) Get(route string, handler http.HandlerFunc) {
	r.route("GET", route, handler)
}

func (r *authenticated) Post(route string, handler http.HandlerFunc) {
	r.route("POST", route, handler)
}

func (r *authenticated) route(method string, route string, handler http.HandlerFunc) {
	s := method + " " + r.basePath + route
	r.log.Debug("Add route for %s", s)
	r.routes[s] = handler
}

func (r *authenticated) Start() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("core/static"))))
	http.HandleFunc("/", r.ServeHTTP)
	r.log.Info("Listening to http://localhost:%d", r.port)
	http.ListenAndServe(fmt.Sprintf(":%d", r.port), nil)
}
