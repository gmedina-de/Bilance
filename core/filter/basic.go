package filter

import (
	"net/http"
)

func Basic(writer http.ResponseWriter, request *http.Request, checker func(username, password string) bool) bool {
	username, password, ok := request.BasicAuth()
	if ok && checker(username, password) {
		return true
	}
	writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	return false
}
