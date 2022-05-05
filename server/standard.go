package server

import (
	"flag"
	"fmt"
	"genuine/filters"
	"genuine/log"
	"genuine/router"
	wd "golang.org/x/net/webdav"
	"net/http"
	"net/url"
	"os"
)

var port = flag.Int("port", 8080, "application port")
var webdav_path = flag.String("webdav_path", "./data", "webdav data path")

type standard struct {
	log     log.Log
	router  router.Router
	filters []filters.Filter
	dav     *wd.Handler
}

func Standard(log log.Log, router router.Router, filters []filters.Filter) Server {
	path := *webdav_path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Critical(err.Error())
		}
	}
	return &standard{log, router, filters, &wd.Handler{
		Prefix:     "/dav/",
		FileSystem: wd.Dir(path),
		LockSystem: wd.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			switch r.Method {
			case "COPY", "MOVE":
				dst := ""
				if u, err := url.Parse(r.Header.Get("Destination")); err == nil {
					dst = u.Path
				}
				log.Debug("WEBDAV %s %s -> %s", r.Method, r.URL.Path, dst)
			default:
				log.Debug("WEBDAV %s %s", r.Method, r.URL.Path)
			}
			if err != nil {
				log.Error(err.Error())
			}
		},
	}}
}

func (s *standard) Serve() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/dav/", s.ServeDAV)
	http.HandleFunc("/", s.ServeHTTP)

	s.log.Info("Starting server http://localhost:%d", *port)
	s.log.Critical(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil).Error())
}

func (s *standard) ServeDAV(writer http.ResponseWriter, request *http.Request) {
	for _, f := range s.filters {
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
	s.dav.ServeHTTP(writer, request)
}

func (s *standard) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handle := true
	for _, f := range s.filters {
		if !f.Filter(writer, request) {
			handle = false
		}
	}
	if handle {
		w := &statusWriter{ResponseWriter: writer, status: 200}
		s.router.Handle(w, request)
		s.log.Debug("%s %s -> %d", request.Method, request.URL, w.status)
	}
}

// in order to know code status
type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
