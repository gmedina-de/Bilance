package controllers

import "genuine/core/controllers"

type index struct {
	*controllers.Base
}

func Index() controllers.Controller {
	return &index{}
}

func (this *index) Routes() map[string]string {
	return map[string]string{
		"/":          "get:Index()",
		"/parameter": "get:Parameter(parameter)",
	}
}

func (this *index) Index() {
}

func (this *index) Parameter(parameter string) {
	this.Data["Title"] = "Parameter id: " + parameter
}
