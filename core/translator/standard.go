package translator

import (
	"fmt"
	"net/http"
	"reflect"
)

type standard struct {
	language     string
	translation  any
	translations map[string]any
}

func Standard() Translator {
	return &standard{translations: make(map[string]any)}
}

func (s *standard) Add(language string, translation any) {
	s.translations[language] = translation
	s.translation = translation
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
	translation, found := s.translations[s.language]
	if found {
		fieldValue := reflect.ValueOf(translation).FieldByName(key)
		if fieldValue.IsValid() {
			return fmt.Sprintf(fieldValue.String(), params...)
		}
	}
	fieldType, _ := reflect.TypeOf(s.translation).FieldByName(key)
	lookup, ok := fieldType.Tag.Lookup("default")
	if ok {
		return fmt.Sprintf(lookup, params...)
	}
	return key
}
