package inject

import (
	"genuine/core/loggers"
	"reflect"
	"strings"
)

var implementations = make(map[reflect.Type][]interface{})
var instanceMap = make(map[reflect.Type]reflect.Value)
var level = 0
var l = loggers.Console()

func Implementations[T any](constructors ...func() T) {
	for _, constructor := range constructors {
		returnType := reflect.ValueOf(constructor).Type().Out(0)
		implementations[returnType] = append(implementations[returnType], constructor)
	}
}

func Inject[T any](pointer T) T {
	pointerValue := reflect.ValueOf(pointer)
	object := pointerValue.Elem()
	l.Debug(strings.Repeat("-", level)+"Injecting %s", pointer)

	for i := 0; i < object.NumField(); i++ {
		field := object.Field(i)
		//
		//lookup, o := object.Type().Field(i).Tag.Lookup("inject")
		//
		//fmt.Println(lookup)
		//fmt.Println(o)

		level++
		instances, ok := instances(field.Type())
		if ok && field.CanSet() {
			if field.Kind() == reflect.Slice {
				field.Set(instances)
			} else {
				field.Set(instances.Index(0))
			}
		}
		level--
	}
	return pointerValue.Interface().(T)
}

func instances(parameterType reflect.Type) (result reflect.Value, ok bool) {
	if parameterType.Kind() == reflect.Slice {
		parameterType = parameterType.Elem()
	}
	instances, found := instanceMap[parameterType]
	if !found {
		constructors, found := implementations[parameterType]
		if !found {
			return reflect.Value{}, false
		}
		instances = reflect.MakeSlice(reflect.SliceOf(parameterType), 0, 0)
		for _, c := range constructors {
			value := reflect.ValueOf(c).Call(nil)[0]
			Inject(value)
			instances = reflect.Append(instances, value)
		}
		instanceMap[parameterType] = instances
	}
	return instances, true
}
