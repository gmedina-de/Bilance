package core

import (
	"fmt"
	"homecloud/core/authenticator"
	"homecloud/core/controller"
	"homecloud/core/database"
	"homecloud/core/log"
	"homecloud/core/repository"
	"homecloud/core/server"
	. "reflect"
	"strings"
)

func init() {
	AddConstructors(
		log.Console,
		database.Gorm,
		authenticator.Basic,
		repository.Users,
		controller.Users,
	)

}

func Init() {
	for k, v := range constructorsMap {
		fmt.Println(k, "constructors: ", len(v))
	}
	fmt.Println("___________________________________")

	srv := inject(server.Authenticated).Interface().(server.Server)
	srv.Start()
}

var (
	constructorsMap = make(map[Type][]interface{})
	instancesMap    = make(map[Type]Value)
	currentLevel    = 0
)

func AddConstructors(constructors ...interface{}) {
	for _, constructor := range constructors {
		constructorReturnType := ValueOf(constructor).Type().Out(0)
		if _, already := constructorsMap[constructorReturnType]; !already {
			constructorsMap[constructorReturnType] = []interface{}{}
		}
		constructorsMap[constructorReturnType] = append(constructorsMap[constructorReturnType], constructor)
	}
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
