package translator

import (
	"net/http"
)

type Translator interface {
	Add(language string, translation any)
	Set(request *http.Request)
	Translate(key string, params ...any) string
}
