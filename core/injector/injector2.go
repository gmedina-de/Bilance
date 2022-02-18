package injector

import (
	log2 "genuine/core/log"
	"github.com/beego/beego/v2/server/web"
	. "reflect"
	"strings"
)

var implementations2 = make(map[Type][]interface{})

func Implementations2[T any](constructors ...func() T) {
	for _, constructor := range constructors {
		returnType := ValueOf(constructor).Type().Out(0)
		implementations2[returnType] = append(implementations[returnType], constructor)
	}
}

var instanceMap = make(map[Type]Value)

var level = 0
var log = log2.Console()

func Inject[T any](strucT T) T {

	object := ValueOf(strucT)
	struc := object.Elem()
	debug("Injecting %s", struc)
	for i := 0; i < struc.NumField(); i++ {
		field := struc.Field(i)
		level++
		if field.CanSet() {
			instances := instances(field.Type())

			if field.Kind() == Slice {
				field.Set(instances)
			} else {
				field.Set(instances.Index(0))
			}

		}
		level--
	}
	return object.Interface().(T)
}

func instances(parameterType Type) Value {
	parameterName := parameterType.Name()
	if parameterType.Kind() == Slice {
		parameterType = parameterType.Elem()
		parameterName = "[]" + parameterType.Name()
	}
	instances, found := instanceMap[parameterType]
	if !found {
		debug("%s wasn't instantiated", parameterName)
		constructors, found := implementations2[parameterType]
		if !found {
			panic("No constructors found for " + parameterName + ", required for dependency injection, provide at least one")
		}
		instances = MakeSlice(SliceOf(parameterType), 0, 0)
		for _, c := range constructors {
			instances = Append(instances, ValueOf(c).Call(nil)[0])
		}
		instanceMap[parameterType] = instances
	} else {
		debug("Parameter %s was already instantiated", parameterName)
	}
	return instances
}

func debug(format string, v ...interface{}) {
	if log != nil && web.BConfig.RunMode == web.DEV {
		log.Debug(strings.Repeat("-", level)+format, v...)
	}
}
