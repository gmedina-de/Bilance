package translator

import (
	"fmt"
	"genuine/core/localizations"
	"html/template"
	"net/http"
)

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
	localization, found := s.localizations[s.language]
	if found {
		translation, found := localization[key]
		if found {
			return fmt.Sprintf(translation, params...)
		}
	}
	localization, found = s.localizations["default"]
	if found {
		translation, found := localization[key]
		if found {
			return fmt.Sprintf(translation, params...)
		}
	}
	return key
}
