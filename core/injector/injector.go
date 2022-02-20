package injector

import "reflect"

type Injector interface {
	Implementations(constructors ...any)
	Inject(constructor any) reflect.Value
}
