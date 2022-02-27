package translator

import (
	"net/http"
)

type Translator interface {
	Set(request *http.Request)
	Translate(key string, params ...any) string
}
