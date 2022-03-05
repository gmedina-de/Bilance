package server

import (
	"fmt"
	"genuine/core/filters"
	"genuine/core/log"
	"genuine/core/server"
	wd "golang.org/x/net/webdav"
	"net/http"
	"net/url"
	"os"
)

type webdav struct {
	log     log.Log
	filters []filters.Filter
}

func Webdav(log log.Log, filters []filters.Filter) server.Server {
	return &webdav{log, filters}
}

func (w *webdav) Serve() {

	path := *webdav_path
	port := *webdav_port

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			w.log.Critical(err.Error())
		}
	}

	fs := wd.Dir(path)
	dav := &wd.Handler{
		FileSystem: fs,
		LockSystem: wd.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			switch r.Method {
			case "COPY", "MOVE":
				dst := ""
				if u, err := url.Parse(r.Header.Get("Destination")); err == nil {
					dst = u.Path
				}
				w.log.Debug("WEBDAV %s %s -> %s", r.Method, r.URL.Path, dst)
			default:
				w.log.Debug("WEBDAV %s %s", r.Method, r.URL.Path)
			}
			if err != nil {
				w.log.Error(err.Error())
			}
		},
	}

	var handler http.Handler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		for _, f := range w.filters {
			if !f.Filter(writer, request) {
				return
			}
		}
		// todo make user-specific folders
		//err := fs.Mkdir(request.Context(), username, 0777)
		//if err != nil && os.IsNotExist(err) {
		//	http.Error(writer, "Error creating user directory", http.StatusInternalServerError)
		//	return
		//}
		dav.ServeHTTP(writer, request)
	})

	w.log.Info("webdav server started dav://localhost:%v", port)
	w.log.Critical(http.ListenAndServe(fmt.Sprintf(":%d", port), handler).Error())
}
