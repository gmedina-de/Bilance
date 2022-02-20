package injector

import (
	"genuine/core/log"
	"reflect"
	"strings"
)

type standard struct {
	constructors map[reflect.Type][]any
	instances    map[reflect.Type]reflect.Value
	level        int
	log          log.Log
}

func Standard() Injector {
	return &standard{
		constructors: make(map[reflect.Type][]any),
		instances:    make(map[reflect.Type]reflect.Value),
		level:        0,
		log:          log.Standard(),
	}
}

func (s *standard) Constructor(constructor any) {
	returnType := reflect.ValueOf(constructor).Type().Out(0)
	s.constructors[returnType] = append(s.constructors[returnType], constructor)
}

func (s *standard) Inject(constructor any) reflect.Value {
	ret := reflect.ValueOf(constructor).Call(nil)[0]
	s.log.Debug(strings.Repeat("  ", s.level)+"Inject %s", ret.Type())

	elem := ret.Elem()
	var value reflect.Value
	if elem.Kind() == reflect.Ptr {
		value = elem.Elem()
	} else {
		value = elem
	}

	for i := 0; i < value.NumField(); i++ {
		s.level++
		field := value.Field(i)
		instances, ok := s.Instances(field.Type())
		if ok && field.CanSet() {
			if field.Kind() == reflect.Slice {
				field.Set(instances)
			} else {
				field.Set(instances.Index(0))
			}
		}
		s.level--
	}

	if elem.Type().Implements(initiableType) {
		elem.Interface().(Initiable).Init()
	}
	return elem
}

func (s *standard) Instances(parameterType reflect.Type) (reflect.Value, bool) {
	if parameterType.Kind() == reflect.Slice {
		parameterType = parameterType.Elem()
	}

	instances, found := s.instances[parameterType]
	if !found {
		constructors, found := s.constructors[parameterType]
		if !found {
			s.log.Warning(strings.Repeat("  ", s.level)+"No constructor found for %s", parameterType.Name())
			return reflect.Value{}, false
		}

		instances = reflect.MakeSlice(reflect.SliceOf(parameterType), 0, 0)
		for _, c := range constructors {

			instances = reflect.Append(instances, s.Inject(c))
		}
		s.instances[parameterType] = instances
	}
	return instances, true
}

func (s *standard) Instance(instance any) {

}
