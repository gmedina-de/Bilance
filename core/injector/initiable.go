package injector

import "reflect"

type Initiable interface {
	Init()
}

var initiableType = reflect.TypeOf((*Initiable)(nil)).Elem()
