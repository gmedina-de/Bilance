package translator

import (
	"fmt"
	"net/http"
)

type standard struct {
	language string
}

func Standard() Translator {
	return &standard{}
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
	localization, found := localizations[s.language]
	if found {
		translation, found := localization[key]
		if found {
			return fmt.Sprintf(translation, params...)
		}
	}
	localization, found = localizations["default"]
	if found {
		translation, found := localization[key]
		if found {
			return fmt.Sprintf(translation, params...)
		}
	}
	return key
}
