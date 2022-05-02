package controllers

type index struct {
}

func Index() Controller {
	return &index{}
}

func (i *index) Routes() Routes {
	return Routes{
		"GET /": i.Index,
	}
}

func (i *index) Index(Request) Response {
	// todo general functions for handlers
	return map[string]any{
		"Template": "index",
	}
}
