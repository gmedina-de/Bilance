package server

import (
	"fmt"
	"homecloud/core/authenticator"
	"homecloud/core/controller"
	"homecloud/core/log"
	"net/http"
	"reflect"
	"strings"
)

type authenticated struct {
	log           log.Log
	authenticator authenticator.Authenticator
	routes        map[string]pair
	port          int
	baseRoute     string
}

type pair struct {
	controller controller.Controller
	action     reflect.Value
}

func Authenticated(log log.Log, authenticator authenticator.Authenticator, controllers []controller.Controller) Server {
	a := &authenticated{log: log, authenticator: authenticator, routes: make(map[string]pair), port: 8080}
	for _, c := range controllers {
		a.baseRoute = reflect.TypeOf(c).Name()
		controllerType := reflect.TypeOf(c)
		name := controllerType.Elem().Name()
		for i := 0; i < controllerType.NumMethod(); i++ {
			method := controllerType.Method(i)
			action := strings.Replace("/"+strings.ToLower(method.Name), "/index", "", 1)
			route := "GET /" + strings.Replace(name, "index", "", 1) + action
			log.Debug("Route %s -> %s", route, method.Func)
			a.routes[route] = pair{c, method.Func}
		}
	}
	return a
}

func (r *authenticated) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.log.Debug("%s %s", request.Method, request.URL)
	if r.authenticator.Authenticate(writer, request) {
		handler, found := r.routes[strings.ToUpper(request.Method)+" "+request.URL.Path]
		if found {
			handler.action.Call([]reflect.Value{
				reflect.ValueOf(handler.controller),
				reflect.ValueOf(writer),
				reflect.ValueOf(request),
			})
		} else {
			fmt.Fprintf(writer, "No route for %s found!", request.URL.Path)
		}
	}
}

func (r *authenticated) Start() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("core/static"))))
	http.HandleFunc("/", r.ServeHTTP)
	r.log.Info("Listening to http://localhost:%d", r.port)
	http.ListenAndServe(fmt.Sprintf(":%d", r.port), nil)
}
