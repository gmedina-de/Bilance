package server

//func Webdav() any {
//
//
//	addr := "0.0.0.0:8081"
//	path := "./data"
//
//	if _, err := os.Stat(path); os.IsNotExist(err) {
//		err := os.Mkdir(path, 0755)
//		if err != nil {
//			log.Fatal(err.Error())
//		}
//	}
//
//	fs := webdav.Dir(path)
//	dav := &webdav.Handler{
//		FileSystem: fs,
//		LockSystem: webdav.NewMemLS(),
//		Logger: func(r *http.Request, err error) {
//			switch r.Method {
//			case "COPY", "MOVE":
//				dst := ""
//				if u, err := url.Parse(r.Header.Get("Destination")); err == nil {
//					dst = u.Path
//				}
//				log.Debug("WEBDAV %s %s -> %s", r.Method, r.URL.Path, dst)
//			default:
//				log.Debug("WEBDAV %s %s", r.Method, r.URL.Path)
//			}
//			if err != nil {
//				log.Error(err.Error())
//			}
//		},
//	}
//
//	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		username, password, ok := r.BasicAuth()
//		if !ok || !auth.Authenticate(username, password) || !strings.HasPrefix(r.URL.Path, "/"+username) {
//			w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
//			http.Error(w, "authorization failed", http.StatusUnauthorized)
//			return
//		}
//
//		err := fs.Mkdir(r.Context(), username, 0777)
//		if err != nil && os.IsNotExist(err) {
//			http.Error(w, "Error creating user directory", http.StatusInternalServerError)
//			return
//		}
//		dav.ServeHTTP(w, r)
//	})
//
//	log.Info("webdav server started http://%v", addr)
//	//http.Handle("/webdav", handler)
//	go func() {
//		log.Fatal(http.ListenAndServe(addr, handler).Error())
//	}()
//	return nil
//}
