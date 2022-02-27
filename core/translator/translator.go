package translator

import (
	"net/http"
)

type Translator interface {
	Add(language string, translation any)
	Translate(language string, key string, params ...any) string
	Lang(request *http.Request) string
}
