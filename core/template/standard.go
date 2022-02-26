package template

import (
	"genuine/core/log"
	"genuine/core/translator"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type standard struct {
	translator translator.Translator
	log        log.Log
	templates  map[string]*template.Template
}

func Standard() Template {
	return &standard{}
}

func (s *standard) Parse(viewsDirectory string) {
	s.templates = make(map[string]*template.Template)

	main := template.New("base.gohtml")
	main.Funcs(template.FuncMap{
		"l10n":     s.translator.Translate,
		"inputs":   inputs,
		"td":       Td,
		"th":       Th,
		"sum":      func(a int, b int) int { return a + b },
		"contains": func(a string, b int64) bool { return strings.Contains(a, strconv.FormatInt(b, 10)) },
	})

	baseFiles, err := filepath.Glob(viewsDirectory + "/base/*.gohtml")
	if err != nil {
		s.log.Error(err.Error())
	}
	files, err := filepath.Glob(viewsDirectory + "/*.gohtml")
	if err != nil {
		s.log.Error(err.Error())
	}
	for _, file := range files {
		tmpl, _ := main.Clone()
		s.templates[filepath.Base(file)] = template.Must(tmpl.ParseFiles(append(baseFiles, file)...))
	}

}

func (s *standard) Render(writer http.ResponseWriter, template string, data map[string]any) {
	if err := s.templates[template].Execute(writer, data); err != nil {
		s.log.Error(err.Error())
	}
}
