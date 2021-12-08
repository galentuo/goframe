package goframe

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	Handle(method string, path string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Router_ struct {
	mux *mux.Router
}

func NewRouter(strictSlash bool) *Router_ {
	r := mux.NewRouter()
	r.StrictSlash(strictSlash)
	return &Router_{r}
}

func (router Router_) Handle(method string, path string, handler http.Handler) {
	router.mux.Handle(path, handler).Methods(method)
}

func (router Router_) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}
