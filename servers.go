package goframe

import (
	"fmt"

	"github.com/galentuo/goframe/logger"
)

// assert Server satisfies Service interface
var _ Service = &Server{}

type Server struct {
	name     string
	logLevel string
}

func (ds Server) Name() string {
	return ds.name
}

func (ds Server) loglevel() string {
	return ds.logLevel
}

func (ds *Server) SetLogLevel(ll logger.LogLevel) {
	ds.logLevel = ll.String()
}

func newService(name string) Server {
	return Server{
		name: name,
	}
}

// assert Server satisfies Service interface
var _ HTTPService = &HTTPServer{}

type HTTPServer struct {
	*Server
	pathPrefix      string
	routeMap        map[string][]EndPoint
	middlewareStack *MiddlewareStack
}

func (dhs *HTTPServer) prefix() string {
	if dhs.pathPrefix == "" {
		return fmt.Sprintf("/%s", dhs.name)
	}
	return dhs.pathPrefix
}

// CustomPrefix replaces the default path prefix by the
// custom one passed in. The routes on the service
// would have the `Service Name` as a default prefix.
func (dhs *HTTPServer) CustomPrefix(prefix string) {
	dhs.pathPrefix = prefix
}

func (dhs *HTTPServer) routes() map[string][]EndPoint {
	return dhs.routeMap
}

func (dhs *HTTPServer) middleware() *MiddlewareStack {
	return dhs.middlewareStack
}

// Use the specified Middleware for the `HTTPService`.
// The specified middleware will be inherited by any calls
// that are made on the HTTPService.
func (dhs *HTTPServer) Use(mw ...MiddlewareFunc) {
	dhs.middlewareStack.Use(mw...)
}

// Route maps a HTTP method request to the path and the specified handler.
func (dhs *HTTPServer) Route(path, httpMethod string, handler HandlerFunction) {
	endpoint := EndPoint{
		httpMethod:      httpMethod,
		handlerFunction: handler,
	}
	var (
		endpoints []EndPoint
		ok        bool
	)
	routes := dhs.routes()
	if endpoints, ok = routes[path]; !ok {
		endpoints = []EndPoint{}
	}

	endpoints = append(endpoints, endpoint)
	dhs.routeMap[path] = endpoints
}

func NewHTTPService(name string) *HTTPServer {
	ds := newService(name)
	routeMap := make(map[string][]EndPoint)
	var mwf []MiddlewareFunc
	mws := MiddlewareStack{
		stack: mwf,
	}
	dhs := HTTPServer{
		Server:          &ds,
		routeMap:        routeMap,
		middlewareStack: &mws,
	}
	return &dhs
}
