package injector

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"log"
	. "reflect"
	"strings"
)

type Injector struct {
	Debug           bool
	ConstructorsMap map[Type][]interface{}
	InstancesMap    map[Type]Value
	CurrentLevel    int
}

func Instance(typ Type) Value {
	instancesMap := Inj.InstancesMap
	fmt.Println(instancesMap)
	value := instancesMap[typ]
	return value
}

func (inj *Injector) Inject(constructor interface{}) Value {
	inj.Debug = web.BConfig.RunMode == web.DEV

	constructorValue := ValueOf(constructor)
	constructorType := constructorValue.Type()
	parameters := make([]Value, constructorType.NumIn())

	inj.debug(strings.Repeat("\t", inj.CurrentLevel), "Instantiating", constructorType)
	for i := 0; i < len(parameters); i++ {
		inj.CurrentLevel++
		instances := inj.instances(constructorType.In(i))
		if constructorType.In(i).Kind() == Slice {
			parameters[i] = instances
		} else {
			parameters[i] = instances.Index(0)
		}
		inj.CurrentLevel--
	}
	return constructorValue.Call(parameters)[0]
}

func (inj *Injector) instances(parameterType Type) Value {
	parameterName := parameterType.Name()
	if parameterType.Kind() == Slice {
		parameterType = parameterType.Elem()
		parameterName = "[]" + parameterType.Name()
	}
	instances, found := inj.InstancesMap[parameterType]
	if !found {
		inj.debug(strings.Repeat("\t", inj.CurrentLevel), parameterName, "wasn't instantiated. ")
		constructors, found := inj.ConstructorsMap[parameterType]
		if !found {
			panic("No constructors found for " + parameterName + ", required for dependency injection, please provide at least one")
		}
		instances = MakeSlice(SliceOf(parameterType), 0, 0)
		for _, c := range constructors {
			instances = Append(instances, inj.Inject(c))
		}
		inj.InstancesMap[parameterType] = instances
	} else {
		inj.debug(strings.Repeat("\t", inj.CurrentLevel), parameterName, "was already instantiated. ")
	}
	return instances
}

func (inj *Injector) debug(a ...any) {
	if inj.Debug {
		log.Println(a...)
	}
}

var Inj = &Injector{
	ConstructorsMap: make(map[Type][]interface{}),
	InstancesMap:    make(map[Type]Value),
	CurrentLevel:    0,
}

func Implementations(constructors ...interface{}) {
	for _, constructor := range constructors {
		constructorReturnType := ValueOf(constructor).Type().Out(0)
		if _, already := Inj.ConstructorsMap[constructorReturnType]; !already {
			Inj.ConstructorsMap[constructorReturnType] = []interface{}{}
		}
		Inj.ConstructorsMap[constructorReturnType] = append(Inj.ConstructorsMap[constructorReturnType], constructor)
	}
}
