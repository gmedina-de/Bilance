package template

import (
	"net/http"
)

type Template interface {
	Parse(viewsDirectory string)
	Render(writer http.ResponseWriter, template string, data map[string]any)
}
