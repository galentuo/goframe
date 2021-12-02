package goframe

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	Handle(method string, path string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type DefaultRouter struct {
	mux *mux.Router
}

func NewRouter() Router {
	return &DefaultRouter{mux.NewRouter()}
}

func (dr DefaultRouter) Handle(method string, path string, handler http.Handler) {
	dr.mux.Handle(path, handler).Methods(method)
}

func (dr DefaultRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dr.mux.ServeHTTP(w, r)
}
