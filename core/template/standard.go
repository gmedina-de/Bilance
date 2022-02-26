package template

import (
	"genuine/core/config"
	"genuine/core/localization"
	"genuine/core/log"
	"genuine/core/navigation"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type standard struct {
	localization localization.Localization
	navigation   navigation.Navigation
	log          log.Log
	templates    map[string]*template.Template
}

const extension = ".gohtml"

func Standard(localization localization.Localization, navigation navigation.Navigation, log log.Log) Template {
	s := &standard{localization, navigation, log, make(map[string]*template.Template)}

	main := template.New("base.gohtml")
	main.Funcs(template.FuncMap{
		"l10n":     s.localization.Translate,
		"inputs":   inputs,
		"td":       Td,
		"th":       Th,
		"sum":      func(a int, b int) int { return a + b },
		"contains": func(a string, b int64) bool { return strings.Contains(a, strconv.FormatInt(b, 10)) },
	})

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
	s.navigation.Handle(data)
	if err := s.templates[template+extension].Execute(writer, data); err != nil {
		s.log.Error(err.Error())
	}
}
