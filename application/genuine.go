package application

import (
	"fmt"
	. "reflect"
	"strings"
)

var (
	constructorsMap = make(map[Type][]interface{})
	instancesMap    = make(map[Type]Value)
	currentLevel    = 0
)

func Genuine(constructors ...interface{}) {
	for _, constructor := range constructors {
		constructorReturnType := ValueOf(constructor).Type().Out(0)
		if _, already := constructorsMap[constructorReturnType]; !already {
			constructorsMap[constructorReturnType] = []interface{}{}
		}
		constructorsMap[constructorReturnType] = append(constructorsMap[constructorReturnType], constructor)
	}

	for k, v := range constructorsMap {
		fmt.Println(k, "constructors: ", len(v))
	}
	fmt.Println("___________________________________")

	a := inject(Application).Interface().(*application)
	fmt.Println("___________________________________")

	for k, v := range instancesMap {
		fmt.Println(k, "instances: ", v.Len())
	}

	a.Run()
}

func inject(constructor interface{}) Value {
	constructorValue := ValueOf(constructor)
	constructorType := constructorValue.Type()
	parameters := make([]Value, constructorType.NumIn())

	fmt.Println(strings.Repeat("\t", currentLevel), "Instantiating", constructorType)
	for i := 0; i < len(parameters); i++ {
		currentLevel++
		parameterType := constructorType.In(i)
		parameterName := parameterType.Name()
		if parameterType.Kind() == Slice {
			parameterType = parameterType.Elem()
			parameterName = "[]" + parameterType.Name()
		}

		instances, found := instancesMap[parameterType]
		if !found {
			fmt.Println(strings.Repeat("\t", currentLevel), parameterName, "wasn't instantiated. ")
			constructors, found := constructorsMap[parameterType]
			if !found {
				panic("No constructors found for " + parameterName + ", required for dependency injection, please provide at least one")
			}
			instances = MakeSlice(SliceOf(parameterType), 0, 0)
			for _, c := range constructors {
				instances = Append(instances, inject(c))
			}
			instancesMap[parameterType] = instances
		} else {
			fmt.Println(strings.Repeat("\t", currentLevel), parameterName, "was already instantiated. ")
		}

		if constructorType.In(i).Kind() == Slice {
			parameters[i] = instances
		} else {
			parameters[i] = instances.Index(0)
		}
		currentLevel--
	}
	return constructorValue.Call(parameters)[0]
}
