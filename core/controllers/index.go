package controllers

type index struct {
	BaseController
}

func Index() Controller {
	return &index{}
}

func (this *index) Route() string {
	this.ViewPath = "core/views"
	return "/"
}

func (this *index) Get() {
	this.TplName = "index.gohtml"
}
