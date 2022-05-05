package router

import (
	controllers3 "genuine/controllers"
	"genuine/database"
	"genuine/decorators"
	"genuine/log"
	"genuine/template"
	"net/http"
	"strings"
)

type standard struct {
	controllers []controllers3.Controller
	decorators  []decorators.Decorator
	template    template.Template
	database    database.Database

	routes map[string]controllers3.Handler
}

func Standard(
	cs []controllers3.Controller,
	decorators []decorators.Decorator,
	template template.Template,
	database database.Database,
	log log.Log,
) Router {
	//todo use regex routing (library?) for more flexibility
	s := &standard{cs, decorators, template, database, make(map[string]controllers3.Handler)}
	for _, c := range s.controllers {
		for k, v := range c.Routes() {
			s.routes[k] = v
			log.Debug("Add route %s", k)
		}
	}
	return s
}

func (s *standard) Handle(w http.ResponseWriter, r *http.Request) {
	action, found := s.routes[strings.ToUpper(r.Method)+" "+r.URL.Path]
	if found {
		request := controllers3.Request{Request: r, ResponseWriter: w}
		response := action(request)
		if response != nil {
			for _, d := range s.decorators {
				d.Decorate(request, response)
			}

			tmpl, render := response["Template"].(string)
			if render {
				response["Database"] = s.database
				s.template.Render(r, w, tmpl, response)
			}
		}
	} else {
		w.WriteHeader(404)
	}
}
