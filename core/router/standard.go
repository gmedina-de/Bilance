package router

import (
	"genuine/core/controllers"
	"genuine/core/log"
	"genuine/core/template"
	"net/http"
	"strings"
)

type standard struct {
	controllers []controllers.Controller
	log         log.Log
	template    template.Template
	routes      map[string]controllers.Handler
}

func Standard(cs []controllers.Controller, log log.Log, template template.Template) Router {
	s := &standard{cs, log, template, make(map[string]controllers.Handler)}
	for _, c := range s.controllers {
		for k, v := range c.Routes() {
			s.routes[k] = v
		}
	}
	return s
}

func (s *standard) Handle(writer http.ResponseWriter, request *http.Request) {
	action, found := s.routes[strings.ToUpper(request.Method)+" "+request.URL.Path]
	if found {
		response := action(request)
		if response == nil {
			response = map[string]any{}
		}

		path, redirect := response["Redirect"]
		if redirect {
			http.Redirect(writer, request, path.(string), http.StatusTemporaryRedirect)
		}

		template, render := response["Template"]
		if render {
			s.template.Render(request, writer, template.(string), response)
		}

	} else {
		writer.WriteHeader(404)
	}
}
