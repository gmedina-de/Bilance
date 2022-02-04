package injector

import "reflect"

type Injector interface {
	AddImplementation(constructor interface{})
	AddInstance(Type reflect.Type, instance interface{})
	Inject(constructor interface{}) reflect.Value
}
