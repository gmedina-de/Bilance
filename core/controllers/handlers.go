package controllers

import (
	"net/http"
)

func Redirect(path string) Handler {
	return func(r Request) Response {
		http.Redirect(r.ResponseWriter, r.Request, path, http.StatusTemporaryRedirect)
		return nil
	}
}
