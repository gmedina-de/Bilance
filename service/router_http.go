package service

import (
	"fmt"
	"net/http"
	"strings"
)

type httpRouter struct {
	log           Log
	authenticator Authenticator
	routes        map[string]Handler
	port          int
}

func HttpRouter(log Log, authenticator Authenticator) Router {
	return &httpRouter{log: log, authenticator: authenticator, routes: make(map[string]Handler), port: 8080}
}

func (r *httpRouter) Get(route string, handler Handler) {
	r.routes["GET "+route] = handler
}

func (r *httpRouter) Post(route string, handler Handler) {
	r.routes["POST "+route] = handler
}

func (r *httpRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

func (r *httpRouter) Start() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", r.ServeHTTP)
	r.log.Info("Listening to http://localhost:%d", r.port)
	http.ListenAndServe(fmt.Sprintf(":%d", r.port), nil)
}
