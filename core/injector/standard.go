package injector

import (
	"genuine/core/log"
	"reflect"
	"strings"
)

const tag = "INJECT"

var implementations = make(map[reflect.Type][]any)
var instanceMap = make(map[reflect.Type]reflect.Value)
var lvl = 0
var l = log.Console() // todo generalize

func Implementations[T any](constructors ...func() T) {
	for _, constructor := range constructors {
		returnType := reflect.ValueOf(constructor).Type().Out(0)
		implementations[returnType] = append(implementations[returnType], constructor)
	}
}

func Inject(constructor any) reflect.Value {
	ret := reflect.ValueOf(constructor).Call(nil)[0]
	l.Debug(tag, level()+"Inject %s", ret.Type())
	elem := ret.Elem()
	var value reflect.Value
	if elem.Kind() == reflect.Ptr {
		value = elem.Elem()
	} else {
		value = elem
	}

	for i := 0; i < value.NumField(); i++ {
		lvl++
		field := value.Field(i)
		instances, ok := instances(field.Type())
		if ok && field.CanSet() {
			if field.Kind() == reflect.Slice {
				field.Set(instances)
			} else {
				field.Set(instances.Index(0))
			}
		}
		lvl--
	}
	if elem.Type().Implements(initiableType) {
		elem.Interface().(Initiable).Init()
	}
	return elem
}

func level() string {
	return strings.Repeat("  ", lvl)
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
			instances = reflect.Append(instances, Inject(c))
		}
		instanceMap[parameterType] = instances
	}
	return instances, true
}
