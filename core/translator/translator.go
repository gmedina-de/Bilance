package translator

import (
	"genuine/core/functions"
	"net/http"
)

type Translator interface {
	functions.Provider
	Set(request *http.Request)
	Translate(key string, params ...any) string
}
