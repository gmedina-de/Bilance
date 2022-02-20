package router

import (
	"genuine/core/controllers"
	"genuine/core/log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type standard struct {
	controllers []controllers.Controller
	log         log.Log
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
	for _, c := range s.controllers {
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
		fun := action.method.Func.Type()
		for i := 0; i < fun.NumIn()-1; i++ {
			queryParam := request.URL.Query().Get(action.parameters[i])
			var value any
			switch fun.In(i).Kind() {
			case reflect.Int:
				value, _ = strconv.Atoi(queryParam)
			case reflect.Bool:
				value, _ = strconv.ParseBool(queryParam)
			default:
				value = queryParam
			}
			args = append(args, reflect.ValueOf(value))
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
			key := strings.ToUpper(m) + " " + route
			s.routes[key] = action{method, controller, s.parameters(v)}
			s.log.Debug("Add route %s", key)
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
