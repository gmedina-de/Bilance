package translator

import (
	"fmt"
	"reflect"
)

type standard struct {
	translations map[string]any
}

func Standard() Translator {
	return &standard{translations: make(map[string]any)}
}

func (s *standard) Translation(language string, translation any) {
	s.translations[language] = translation
}

func (s *standard) Translate(language string, key string, params ...any) string {
	translation, found := s.translations[language]
	if found {
		fieldValue := reflect.ValueOf(translation).FieldByName(key)
		if fieldValue.IsValid() {
			return fmt.Sprintf(fieldValue.String(), params...)
		}
		fieldType, _ := reflect.TypeOf(translation).FieldByName(key)
		lookup, ok := fieldType.Tag.Lookup("default")
		if ok {
			return fmt.Sprintf(lookup, params...)
		}
	}
	return key
}
