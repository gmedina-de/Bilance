package models

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"reflect"
	"strings"

	"github.com/jinzhu/inflection"
)

func Name(model any) string {
	return strings.ToLower(RealValueOf(model).Type().Name())
}

func Plural(model any) string {
	return inflection.Plural(Name(model))
}

func RealValueOf(v interface{}) reflect.Value {
	rv := reflect.ValueOf(v)
	if rv.Type().Kind() == reflect.Ptr && rv.IsNil() {
		rv = reflect.New(rv.Type().Elem()).Elem()
	}
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	return rv
}

func Serialize(T any) string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(T)
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func Deserialize[T any](str string, dest *T) {
	by, _ := base64.StdEncoding.DecodeString(str)
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	_ = d.Decode(&dest)
}
