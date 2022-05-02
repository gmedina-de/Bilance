package controllers

type files struct {
}

func Files() Controller {
	return &files{}
}

func (f *files) Routes() map[string]Handler {
	return map[string]Handler{
		"GET /files": f.Index,
	}
}

func (f *files) Index(Request) Response {
	return Response{
		"Template": "files",
	}
}
