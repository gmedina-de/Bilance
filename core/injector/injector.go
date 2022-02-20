package injector

import (
	"reflect"
)

type Injector interface {
	Constructor(constructor any)
	Inject(constructor any) reflect.Value
	Instances(parameterType reflect.Type) (reflect.Value, bool)
	Instance(instance any)
}
