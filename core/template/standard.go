package template

import (
	"genuine/core/config"
	"genuine/core/log"
	"genuine/core/translator"
	"html/template"
	"net/http"
	"path/filepath"
)

type standard struct {
	localization translator.Translator
	log          log.Log
	templates    map[string]*template.Template
}

const extension = ".gohtml"

func Standard(localization translator.Translator, log log.Log) Template {
	s := &standard{localization, log, make(map[string]*template.Template)}
	main := template.New("base.gohtml")
	AddFunc("l10n", s.localization.Translate)
	main.Funcs(funcMap)

	baseFiles, err := filepath.Glob(config.ViewDirectory() + "/base/*" + extension)
	if err != nil {
		s.log.Error(err.Error())
	}
	files, err := filepath.Glob(config.ViewDirectory() + "/*" + extension)
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
	data["Lang"] = s.localization.Lang(request)
	data["Path"] = request.URL.Path
	if err := s.templates[template+extension].Execute(writer, data); err != nil {
		s.log.Error(err.Error())
	}
}
