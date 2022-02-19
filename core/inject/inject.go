package inject

import (
	"genuine/core/log"
	"reflect"
	"strings"
)

var implementations = make(map[reflect.Type][]any)
var instanceMap = make(map[reflect.Type]reflect.Value)
var level = 0
var l = log.Console()
var initiableType = reflect.TypeOf((*Initiable)(nil)).Elem()

func Implementations[T any](constructors ...func() T) {
	for _, constructor := range constructors {
		returnType := reflect.ValueOf(constructor).Type().Out(0)
		implementations[returnType] = append(implementations[returnType], constructor)
	}
}

func Call(constructor any) reflect.Value {
	elem := reflect.ValueOf(constructor).Call(nil)[0].Elem()
	var value reflect.Value
	if elem.Kind() == reflect.Ptr {
		value = elem.Elem()
	} else {
		value = elem
	}
	l.Debug(strings.Repeat("-", level)+"Injecting %s", value)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
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
	if elem.Type().Implements(initiableType) {
		elem.Interface().(Initiable).Init()
	}
	return elem
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
			instances = reflect.Append(instances, Call(c))
		}
		instanceMap[parameterType] = instances
	}
	return instances, true
}
