package router

import (
	"genuine/core/controllers"
	"net/http"
)

func Redirect(path string) controllers.Handler {
	return func(r controllers.Request) controllers.Response {
		http.Redirect(r.ResponseWriter, r.Request, path, http.StatusTemporaryRedirect)
		return nil
	}
}
