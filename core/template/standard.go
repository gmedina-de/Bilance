package template

import (
	"genuine/core/log"
	"genuine/core/translator"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type standard struct {
	translator translator.Translator
	log        log.Log
	template   *template.Template
	templates  []string
}

func Standard() Template {
	return &standard{}
}

func (s *standard) Templates(base string, templates ...string) {
	s.templates = append(templates, base)
	s.template = template.New(path.Base(base))
	s.template.Funcs(template.FuncMap{
		"l10n":     s.translator.Translate,
		"td":       Td,
		"th":       Th,
		"sum":      func(a int, b int) int { return a + b },
		"contains": func(a string, b int64) bool { return strings.Contains(a, strconv.FormatInt(b, 10)) },
	})
}

func (s *standard) Render(writer http.ResponseWriter, template string, data map[string]any) {
	s.template.ParseFiles(append(s.templates, template)...)
	err := s.template.Execute(writer, data)
	if err != nil {
		s.log.Error(err.Error())
	}
}
