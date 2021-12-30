package goframe

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router registers routes to be matched and dispatches a handler.
type Router interface {
	Handle(method string, path string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type defaultRouter struct {
	mux *mux.Router
}

// NewRouter returns an instance of an implementation of
// Router interface.
func NewRouter(strictSlash bool) *defaultRouter {
	r := mux.NewRouter()
	r.StrictSlash(strictSlash)
	return &defaultRouter{r}
}

// Handle registers a new route with a matcher for the URL path.
func (router *defaultRouter) Handle(method string, path string, handler http.Handler) {
	router.mux.Handle(path, handler).Methods(method)
}

// ServeHTTP should write reply headers and data to the ResponseWriter
// and then return. Returning signals that the request is finished; it
// is not valid to use the ResponseWriter or read from the
// Request.Body after or concurrently with the completion of the
// ServeHTTP call.
func (router *defaultRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}
