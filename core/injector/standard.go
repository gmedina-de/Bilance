package injector

import (
	"genuine/core/log"
	"reflect"
	"strings"
	"unsafe"
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

func (s *standard) Implementation(constructor any) {
	returnType := reflect.ValueOf(constructor).Type().Out(0)
	s.constructors[returnType] = append(s.constructors[returnType], constructor)
	s.log.Debug("Added constructor for type %s", returnType)
}

func (s *standard) Inject(constructor any) reflect.Value {
	value := reflect.ValueOf(constructor).Call(nil)[0]
	s.log.Debug(strings.Repeat("  ", s.level)+"Inject %s", value.Type())

	elem := value
	if elem.Kind() == reflect.Interface {
		elem = elem.Elem()
	}
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}

	for i := 0; i < elem.NumField(); i++ {
		s.level++
		field := elem.Field(i)
		instances, ok := s.Instances(field.Type())
		if ok {
			s.log.Debug(strings.Repeat("  ", s.level)+"Set %s", field.Type())

			// access unexported fields using unsafe Pointer
			field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()

			if field.Kind() == reflect.Slice {
				field.Set(instances)
			} else {
				field.Set(instances.Index(0))
			}
		}
		s.level--
	}

	return value
}

func (s *standard) Instances(parameterType reflect.Type) (reflect.Value, bool) {
	if parameterType.Kind() == reflect.Slice {
		parameterType = parameterType.Elem()
	}

	instances, found := s.instances[parameterType]
	if !found {
		constructors, found := s.constructors[parameterType]
		if !found {
			s.log.Debug(strings.Repeat("  ", s.level)+"No constructor found for %s, ignoring", parameterType.Name())
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
