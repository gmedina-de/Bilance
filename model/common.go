package model

import (
	"reflect"
	"strings"
)

func ModelNamePlural(model interface{}) string {
	of := reflect.TypeOf(model).Elem()
	name := of.Name()
	lower := strings.ToLower(name)
	if strings.HasSuffix(lower, "y") {
		return lower[:len(lower)-1] + "ies"
	}
	return lower + "s"
}
