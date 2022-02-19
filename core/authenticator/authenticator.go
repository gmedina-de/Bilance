package authenticator

import "net/http"

type Authenticator interface {
	Authenticate(writer http.ResponseWriter, request *http.Request) bool
}
