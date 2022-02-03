package injector

import "reflect"

type Injector interface {
	Add(constructors ...interface{})
	Inject(constructor interface{}) reflect.Value
}
