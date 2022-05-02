package template

import (
	"flag"
	functions2 "genuine/functions"
	"genuine/log"
	"html/template"
	"net/http"
	"path/filepath"
)

type standard struct {
	translator functions2.Translator
	log        log.Log
	templates  map[string]*template.Template
}

const extension = ".gohtml"

var views = flag.String("views", "views", "directory where views are stored")

func Standard(translator functions2.Translator, providers []functions2.Provider, log log.Log) Template {

	s := &standard{translator, log, make(map[string]*template.Template)}
	main := template.New("base.gohtml")

	var funcMap = make(template.FuncMap)
	for _, p := range providers {
		funcMap = s.joinFuncMaps(p.GetFuncMap(), funcMap)
	}
	funcMap = s.joinFuncMaps(translator.GetFuncMap(), funcMap)
	main.Funcs(funcMap)

	baseFiles, err := filepath.Glob(*views + "/include/*" + extension)
	if err != nil {
		s.log.Error(err.Error())
	}
	files, err := filepath.Glob(*views + "/*" + extension)
	if err != nil {
		s.log.Error(err.Error())
	}
	for _, file := range files {
		tmpl, _ := main.Clone()
		s.templates[filepath.Base(file)] = template.Must(tmpl.ParseFiles(append(baseFiles, file)...))
	}

	return s
}

func (s *standard) Render(request *http.Request, writer http.ResponseWriter, template string, data map[string]any) {
	data["Path"] = request.URL.Path
	s.translator.Set(request)
	templateName := template + extension
	tmpl, found := s.templates[templateName]
	if found {
		if err := tmpl.Execute(writer, data); err != nil {
			s.log.Error(err.Error())
		}
	} else {
		s.log.Error("Template %s not found", templateName)
	}
}

func (s *standard) joinFuncMaps(m1, m2 template.FuncMap) template.FuncMap {
	for k, v := range m1 {
		m2[k] = v
	}
	return m2
}
