package core

import (
	"github.com/beego/beego/v2/server/web"
	. "homecloud/core/controllers"
	"homecloud/core/database"
	log2 "homecloud/core/log"
	"homecloud/core/repositories"
	"homecloud/core/template"
	"log"
	. "reflect"
	"strings"
)

func init() {
	web.AddViewPath("core/views")

	template.AddNavigation("home", "home").Path = "/"
	Implementations(
		log2.Console,
		database.Orm,
		repositories.Users,
		Index,
		Users,
	)

	web.BConfig.AppName = "HomeCloud"
	web.BConfig.Listen.HTTPAddr = "0.0.0.0"
	web.BConfig.Listen.HTTPPort = 8080
	web.BConfig.RunMode = web.DEV
	web.BConfig.WebConfig.AutoRender = true
	web.BConfig.RecoverPanic = false
}

func Init() {

	users := USER{web.Controller{}, database.Orm()}
	web.Router("/task/", users, "get:List")
	web.Run()

	//inj.Inject(homecloud)
}

func homecloud(controllers []Controller) {
	// adds settings to the end of navigation
	template.AddNavigation("settings", "settings").
		WithChild("users", "users").
		Path = "/core/users"

	//for _, c := range controllers {
	//web.AutoPrefix("/"+strings.Replace(PackageName(c), "core", "settings", 1), c)
	//web.AutoPrefix("/"+PackageName(c), c)
	//}

	web.Run()
}

var inj = &injector{
	constructorsMap: make(map[Type][]interface{}),
	instancesMap:    make(map[Type]Value),
	currentLevel:    0,
}

const injectorDebug = true

type injector struct {
	constructorsMap map[Type][]interface{}
	instancesMap    map[Type]Value
	currentLevel    int
}

func Implementations(constructors ...interface{}) {
	for _, constructor := range constructors {
		constructorReturnType := ValueOf(constructor).Type().Out(0)
		if _, already := inj.constructorsMap[constructorReturnType]; !already {
			inj.constructorsMap[constructorReturnType] = []interface{}{}
		}
		inj.constructorsMap[constructorReturnType] = append(inj.constructorsMap[constructorReturnType], constructor)
	}
}

func (inj *injector) Inject(constructor interface{}) Value {
	constructorValue := ValueOf(constructor)
	constructorType := constructorValue.Type()
	parameters := make([]Value, constructorType.NumIn())

	inj.debug(strings.Repeat("\t", inj.currentLevel), "Instantiating", constructorType)
	for i := 0; i < len(parameters); i++ {
		inj.currentLevel++
		instances := inj.instances(constructorType.In(i))
		if constructorType.In(i).Kind() == Slice {
			parameters[i] = instances
		} else {
			parameters[i] = instances.Index(0)
		}
		inj.currentLevel--
	}
	return constructorValue.Call(parameters)[0]
}

func (inj *injector) instances(parameterType Type) Value {
	parameterName := parameterType.Name()
	if parameterType.Kind() == Slice {
		parameterType = parameterType.Elem()
		parameterName = "[]" + parameterType.Name()
	}
	instances, found := inj.instancesMap[parameterType]
	if !found {
		inj.debug(strings.Repeat("\t", inj.currentLevel), parameterName, "wasn't instantiated. ")
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
		inj.debug(strings.Repeat("\t", inj.currentLevel), parameterName, "was already instantiated. ")
	}
	return instances
}

func (inj *injector) debug(a ...any) {
	if injectorDebug {
		log.Println(a...)
	}
}
