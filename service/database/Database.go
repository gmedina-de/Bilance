package database

import (
	"reflect"
)

type Database interface {
	Create(model interface{})
	RetrieveAll(model interface{})
}

func modelName(model interface{}) string {
	return reflect.TypeOf(model).String()
}
