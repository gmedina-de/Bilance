package template

import (
	"net/http"
)

type Template interface {
	Templates(base string, templates ...string)
	Render(writer http.ResponseWriter, template string, data map[string]any)
}
