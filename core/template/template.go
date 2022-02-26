package template

import (
	"net/http"
)

type Template interface {
	Render(request *http.Request, writer http.ResponseWriter, template string, data map[string]any)
}
