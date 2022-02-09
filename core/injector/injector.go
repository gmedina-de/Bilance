package injector

import (
	"genuine/core/log"
	"github.com/beego/beego/v2/server/web"
	. "reflect"
	"strings"
)

var implementations = make(map[Type][]interface{})

func Implementations(constructors ...interface{}) {
	for _, constructor := range constructors {
		returnType := ValueOf(constructor).Type().Out(0)
		implementations[returnType] = append(implementations[returnType], constructor)
	}
}

type injector struct {
	instanceMap map[Type]Value
	level       int
	log         log.Log
}

func Injector(constructor interface{}) *injector {
	i := &injector{
		instanceMap: make(map[Type]Value),
		level:       0,
	}

	// log must be instantiated at first place, because injector depends on it
	i.log = i.instances(TypeOf((*log.Log)(nil)).Elem()).Interface().([]log.Log)[0].(log.Log)
	i.construct(constructor)
	return i
}

func (inj *injector) construct(constructor interface{}) Value {
	constructorValue := ValueOf(constructor)
	constructorType := constructorValue.Type()
	parameters := make([]Value, constructorType.NumIn())

	inj.debug("Instantiating %s", constructorValue)
	for i := 0; i < len(parameters); i++ {
		inj.level++
		instances := inj.instances(constructorType.In(i))
		if constructorType.In(i).Kind() == Slice {
			parameters[i] = instances
		} else {
			parameters[i] = instances.Index(0)
		}
		inj.level--
	}
	return constructorValue.Call(parameters)[0]
}

func (inj *injector) instances(parameterType Type) Value {
	parameterName := parameterType.Name()
	if parameterType.Kind() == Slice {
		parameterType = parameterType.Elem()
		parameterName = "[]" + parameterType.Name()
	}
	instances, found := inj.instanceMap[parameterType]
	if !found {
		inj.debug("%s wasn't instantiated", parameterName)
		constructors, found := implementations[parameterType]
		if !found {
			panic("No constructors found for " + parameterName + ", required for dependency injection, provide at least one")
		}
		instances = MakeSlice(SliceOf(parameterType), 0, 0)
		for _, c := range constructors {
			instances = Append(instances, inj.construct(c))
		}
		inj.instanceMap[parameterType] = instances
	} else {
		inj.debug("Parameter %s was already instantiated", parameterName)
	}
	return instances
}

func (inj *injector) debug(format string, v ...interface{}) {
	if inj.log != nil && web.BConfig.RunMode == web.DEV {
		inj.log.Debug(strings.Repeat("-", inj.level)+format, v...)
	}
}
