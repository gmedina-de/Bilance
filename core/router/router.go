package router

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
	"reflect"
	"strings"
)

func Add(controller web.ControllerInterface, route string, mappingMethods ...string) {
	methods := parseMappingMethods(mappingMethods)
	of := reflect.TypeOf(controller).Elem()
	key := of.PkgPath() + ":" + of.Name()
	for k, v := range methods {
		name := k[:strings.Index(k, "(")]
		allowMethods := strings.Split(v, ",")
		parameters := strings.Split(k[strings.Index(k, "(")+1:strings.Index(k, ")")], ",")
		var params []*param.MethodParam
		for _, p := range parameters {
			if p != "" {
				split := strings.Split(p, " ")
				var options []param.MethodParamOption
				for i, s := range split {
					if i == 0 {
						continue
					}
					switch s {
					case "required":
						options = append(options, param.IsRequired)
					case "header":
						options = append(options, param.InHeader)
					case "path":
						options = append(options, param.InPath)
					case "body":
						options = append(options, param.InBody)
					}
				}
				params = append(params, param.New(split[0], options...))
			}
		}
		comments := web.ControllerComments{
			Method:           name,
			Router:           route,
			AllowHTTPMethods: allowMethods,
			MethodParams:     params,
			Filters:          nil,
			Params:           nil,
		}
		web.GlobalControllerRouter[key] = append(web.GlobalControllerRouter[key], comments)
	}
}

func parseMappingMethods(mappingMethods []string) map[string]string {
	methods := make(map[string]string)
	semi := strings.Split(mappingMethods[0], ";")
	for _, v := range semi {
		colon := strings.Split(v, ":")
		methods[colon[1]] = colon[0]
	}
	return methods
}
