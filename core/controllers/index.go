package controllers

type index struct {
	Base
}

func Index() Controller {
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
