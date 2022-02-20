package injector

import "reflect"

type Injector interface {
	Implementation(constructor any)
	Inject(constructor any) reflect.Value
	Instances(parameterType reflect.Type) (reflect.Value, bool)
}
