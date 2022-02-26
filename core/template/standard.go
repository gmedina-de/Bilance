package template

import (
	"genuine/core/config"
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

const extension = ".gohtml"

func Standard(translator translator.Translator, log log.Log) Template {
	s := &standard{translator, log, make(map[string]*template.Template)}

	main := template.New("base.gohtml")
	main.Funcs(template.FuncMap{
		"l10n":     s.translator.Translate,
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

	lang := "en-US"
	al := request.Header.Get("Accept-Language")
	if len(al) > 4 {
		lang = al[:5]
	}
	data["Lang"] = lang

	path := request.URL.Path
	data["Path"] = path

	currentNavigation1 := GetCurrentNavigation(path, Navigation)
	data["Navigation1"] = Navigation
	data["CurrentNavigation1"] = currentNavigation1
	data["CurrentNavigation1Index"] = GetCurrentNavigationIndex(path, Navigation)
	if currentNavigation1 != nil {
		data["Navigation2"] = currentNavigation1.SubMenu
		data["CurrentNavigation2"] = GetCurrentNavigation(path, currentNavigation1.SubMenu)
		data["CurrentNavigation2Index"] = GetCurrentNavigationIndex(path, currentNavigation1.SubMenu)
	}

	if err := s.templates[template+extension].Execute(writer, data); err != nil {
		s.log.Error(err.Error())
	}
}
