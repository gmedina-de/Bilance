package controllers

import "github.com/beego/beego/v2/server/web"

type index struct {
	BaseController
}

func Index() Controller {
	i := &index{}
	web.Router("/", i, "get:Index")
	return i
}

func (this *index) Index() {
}
