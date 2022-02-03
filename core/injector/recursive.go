package injector

import (
	"fmt"
	. "reflect"
	"strings"
)

type recursive struct {
	constructorsMap map[Type][]interface{}
	instancesMap    map[Type]Value
	currentLevel    int
}

func Recursive() *recursive {
	return &recursive{
		constructorsMap: make(map[Type][]interface{}),
		instancesMap:    make(map[Type]Value),
		currentLevel:    0,
	}
}

func (inj *recursive) Add(constructors ...interface{}) {
	for _, constructor := range constructors {
		constructorReturnType := ValueOf(constructor).Type().Out(0)
		if _, already := inj.constructorsMap[constructorReturnType]; !already {
			inj.constructorsMap[constructorReturnType] = []interface{}{}
		}
		inj.constructorsMap[constructorReturnType] = append(inj.constructorsMap[constructorReturnType], constructor)
	}
}

func (inj *recursive) Inject(constructor interface{}) Value {
	constructorValue := ValueOf(constructor)
	constructorType := constructorValue.Type()
	parameters := make([]Value, constructorType.NumIn())

	fmt.Println(strings.Repeat("\t", inj.currentLevel), "Instantiating", constructorType)
	for i := 0; i < len(parameters); i++ {
		inj.currentLevel++
		parameterType := constructorType.In(i)
		parameterName := parameterType.Name()
		if parameterType.Kind() == Slice {
			parameterType = parameterType.Elem()
			parameterName = "[]" + parameterType.Name()
		}

		instances, found := inj.instancesMap[parameterType]
		if !found {
			fmt.Println(strings.Repeat("\t", inj.currentLevel), parameterName, "wasn't instantiated. ")
			constructors, found := inj.constructorsMap[parameterType]
			if !found {
				panic("No constructors found for " + parameterName + ", required for dependency injection, please provide at least one")
			}
			instances = MakeSlice(SliceOf(parameterType), 0, 0)
			for _, c := range constructors {
				instances = Append(instances, inj.Inject(c))
			}
			inj.instancesMap[parameterType] = instances
		} else {
			fmt.Println(strings.Repeat("\t", inj.currentLevel), parameterName, "was already instantiated. ")
		}

		if constructorType.In(i).Kind() == Slice {
			parameters[i] = instances
		} else {
			parameters[i] = instances.Index(0)
		}
		inj.currentLevel--
	}
	return constructorValue.Call(parameters)[0]
}
