package functions

import (
	"fmt"
	"genuine/localizations"
	"html/template"
	"net/http"
	"strings"
)

type Translator interface {
	Provider
	Set(request *http.Request)
	Translate(key string, params ...any) string
}

type standard struct {
	language      string
	localizations map[string]localizations.Localization
}

func Standard(provider localizations.Provider) Translator {
	return &standard{localizations: provider.GetLocalizations()}
}

func (s *standard) GetFuncMap() template.FuncMap {
	return map[string]any{
		"l10n": s.Translate,
	}
}

func (s *standard) Set(request *http.Request) {
	lang := "en-US"
	al := request.Header.Get("Accept-Language")
	if len(al) > 4 {
		lang = al[:5]
	}
	s.language = lang
}

func (s *standard) Translate(key string, params ...any) string {
	lowerKey := strings.ToLower(key)
	localization, found := s.localizations[s.language]
	if found {
		translation, found := localization[lowerKey]
		if found {
			return fmt.Sprintf(translation, params...)
		}
	}
	localization, found = s.localizations["default"]
	if found {
		translation, found := localization[lowerKey]
		if found {
			return fmt.Sprintf(translation, params...)
		}
	}
	return key
}
