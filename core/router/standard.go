package router

import (
	"genuine/app/database"
	"genuine/core/controllers"
	"genuine/core/decorators"
	"genuine/core/filters"
	"genuine/core/log"
	"genuine/core/template"
	"net/http"
	"strings"
)

type standard struct {
	controllers []controllers.Controller
	filters     []filters.Filter
	decorators  []decorators.Decorator
	template    template.Template
	database    database.Database

	routes map[string]controllers.Handler
}

func Standard(
	cs []controllers.Controller,
	filters []filters.Filter,
	decorators []decorators.Decorator,
	template template.Template,
	database database.Database,
	log log.Log,
) Router {
	s := &standard{cs, filters, decorators, template, database, make(map[string]controllers.Handler)}
	for _, c := range s.controllers {
		for k, v := range c.Routes() {
			s.routes[k] = v
			log.Debug("Add route %s", k)
		}
	}
	return s
}

func (s *standard) Handle(w http.ResponseWriter, r *http.Request) {
	handle := true
	for _, f := range s.filters {
		if !f.Filter(w, r) {
			handle = false
		}
	}
	if handle {
		action, found := s.routes[strings.ToUpper(r.Method)+" "+r.URL.Path]
		if found {
			request := controllers.Request{Request: r, ResponseWriter: w}
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
}
