package filter

import "net/http"

type Filter interface {
	Filter(writer http.ResponseWriter, request *http.Request) bool
}
