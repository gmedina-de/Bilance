package injector

import "reflect"

type Injector interface {
	Add(constructor interface{})
	Inject(constructor interface{}) reflect.Value
}
