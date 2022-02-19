package router

import (
	"genuine/core/controllers"
	"net/http"
	"reflect"
	"strings"
)

type standard struct {
	Controllers []controllers.Controller
	routes      map[string]action
}

type action struct {
	method     reflect.Method
	controller controllers.Controller
	parameters []string
}

func Standard() Router {
	return &standard{routes: make(map[string]action)}
}

func (s *standard) Init() {
	for _, c := range s.Controllers {
		for k, v := range c.Routes() {
			s.addRoute(c, k, v)
		}
	}
}

func (s *standard) Handle(writer http.ResponseWriter, request *http.Request) {
	action, found := s.routes[strings.ToUpper(request.Method)+" "+request.URL.Path]
	if found {
		action.controller.Before(request, writer, reflect.TypeOf(action.controller).Elem().Name())
		args := []reflect.Value{reflect.ValueOf(action.controller)}
		for i := 0; i < action.method.Func.Type().NumIn()-1; i++ {
			args = append(args, reflect.ValueOf(request.URL.Query().Get(action.parameters[i])))
		}
		action.method.Func.Call(args)
		action.controller.After()

	} else {
		writer.WriteHeader(404)
	}
}

func (s *standard) addRoute(controller controllers.Controller, route string, mappingMethods ...string) {
	semicolons := make(map[string]string)
	for _, v := range strings.Split(mappingMethods[0], ";") {
		colon := strings.Split(v, ":")
		semicolons[colon[0]] = colon[1]
	}
	for k, v := range semicolons {
		for _, m := range strings.Split(k, ",") {
			name := v[:strings.Index(v, "(")]
			method, _ := reflect.TypeOf(controller).MethodByName(name)
			s.routes[strings.ToUpper(m)+" "+route] = action{method, controller, s.parameters(v)}
		}
	}
}

func (s *standard) parameters(v string) []string {
	var parameters []string
	commas := strings.Split(v[strings.Index(v, "(")+1:strings.Index(v, ")")], ",")
	for _, p := range commas {
		parameters = append(parameters, p)
	}
	return parameters
}
