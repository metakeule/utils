package http

import (
	"net/http"
)

// mounts the given handler (often a *http.ServeMux) under the
// given path, i.e. all routes beginning with path+"/" with
// be routed to the handler and the path is being stripped away
// so that it appears to the handler as if it would handle the
// root path
func MountToMux(path string, handler http.Handler, mux *http.ServeMux) {
	mux.Handle(path+"/", http.StripPrefix(path, handler))
}

// calls MountToMux  for the http.DefaultServeMux
func Mount(path string, handler http.Handler) {
	MountToMux(path, handler, http.DefaultServeMux)
}
