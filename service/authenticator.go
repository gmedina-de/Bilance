package service

import (
	"net/http"
)

type Authenticator interface {
	Authenticate(w http.ResponseWriter, r *http.Request) bool
}
