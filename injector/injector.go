package injector

import (
	"genuine/functions"
	"genuine/log"
	"genuine/router"
	"genuine/server"
	"genuine/template"
	. "log"
	"reflect"
	"strings"
)

func init() {
	Provide(log.Standard)
	Provide(server.Standard)
	Provide(router.Standard)
	Provide(functions.Standard)
	Provide(template.Standard)
}

// DEPENDENCY INJECTION
var (
	constructorMap = make(map[reflect.Type][]any)
	instanceMap    = make(map[reflect.Type]reflect.Value)
	level          = 0
)

func Provide(constructors ...any) {
	for _, constructor := range constructors {
		value := reflect.ValueOf(constructor)
		if value.Kind() != reflect.Func {
			panic(value.String() + " is not a constructor")
		}

		returnType := value.Type().Out(0)
		constructorMap[returnType] = append(constructorMap[returnType], constructor)
	}
}

func Invoke(constructor any) reflect.Value {
	constructorValue := reflect.ValueOf(constructor)
	constructorType := constructorValue.Type()
	parameters := make([]reflect.Value, constructorType.NumIn())

	debug("Resolving dependencies for %s", constructorType)
	for i := 0; i < len(parameters); i++ {
		level++
		parameterType := constructorType.In(i)
		parameterName := parameterType.Name()
		if parameterType.Kind() == reflect.Slice {
			parameterType = parameterType.Elem()
			parameterName = "[]" + parameterType.Name()
		}

		instances, instancesFound := instanceMap[parameterType]
		if !instancesFound {
			debug("%s wasn't instantiated", parameterName)
			constructors, constructorsFound := constructorMap[parameterType]

			if !constructorsFound {
				panic("No constructor found for " + parameterName)
			}
			instances = reflect.MakeSlice(reflect.SliceOf(parameterType), 0, 10)
			for _, c := range constructors {
				instances = reflect.Append(instances, Invoke(c))
			}

			instanceMap[parameterType] = instances
		} else {
			debug("%s was already instantiated. ", parameterName)
		}

		if constructorType.In(i).Kind() == reflect.Slice {
			parameters[i] = instances
		} else {
			parameters[i] = instances.Index(instances.Len() - 1)
		}
		level--
	}
	debug("Invoking %s", constructorType)
	return constructorValue.Call(parameters)[0]
}

func debug(format string, v ...any) {
	if *log.LogLevel == int(log.Debug) {
		Printf(strings.Repeat("  ", level)+format, v...)
	}
}
