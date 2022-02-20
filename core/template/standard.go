package template

import (
	"genuine/core/log"
	"genuine/core/translator"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
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

func (s *standard) Parse(directory string) {
	s.template = template.New("")
	s.template.Funcs(template.FuncMap{
		"l10n":     s.translator.Translate,
		"inputs":   inputs,
		"td":       Td,
		"th":       Th,
		"sum":      func(a int, b int) int { return a + b },
		"contains": func(a string, b int64) bool { return strings.Contains(a, strconv.FormatInt(b, 10)) },
	})

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".gohtml") {
			_, err = s.template.ParseFiles(path)
			if err != nil {
				s.log.Error(err.Error())
			}
		}
		return err
	})
	if err != nil {
		s.log.Error(err.Error())
	}
}

func (s *standard) Render(writer http.ResponseWriter, template string, data map[string]any) {
	err := s.template.ExecuteTemplate(writer, template, data)
	if err != nil {
		s.log.Error(err.Error())
	}
}
