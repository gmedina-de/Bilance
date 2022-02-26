package localization

import (
	"fmt"
	"net/http"
	"reflect"
)

type standard struct {
	translation  any
	translations map[string]any
}

func Standard() Localization {
	return &standard{translations: make(map[string]any)}
}

func (s *standard) Add(language string, translation any) {
	s.translations[language] = translation
	s.translation = translation
}

func (s *standard) Translate(language string, key string, params ...any) string {
	translation, found := s.translations[language]
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

func (s *standard) Lang(request *http.Request) string {
	lang := "en-US"
	al := request.Header.Get("Accept-Language")
	if len(al) > 4 {
		lang = al[:5]
	}
	return lang
}
