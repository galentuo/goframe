package goframe

import (
	"fmt"

	"github.com/galentuo/goframe/logger"
)

// assert Service_ satisfies Service interface
var _ Service = &Service_{}

type Service_ struct {
	name     string
	logLevel string
}

func (ds Service_) Name() string {
	return ds.name
}

func (ds Service_) loglevel() string {
	return ds.logLevel
}

func (ds *Service_) SetLogLevel(ll logger.LogLevel) {
	ds.logLevel = ll.String()
}

func newService(name string) Service_ {
	return Service_{
		name: name,
	}
}

// assert Service_ satisfies Service interface
var _ HTTPService = &HTTPService_{}

type HTTPService_ struct {
	*Service_
	pathPrefix      string
	routeMap        map[string][]EndPoint
	middlewareStack *MiddlewareStack
}

func (dhs *HTTPService_) prefix() string {
	if dhs.pathPrefix == "" {
		return fmt.Sprintf("/%s", dhs.name)
	}
	return dhs.pathPrefix
}

// CustomPrefix replaces the default path prefix by the
// custom one passed in. The routes on the service
// would have the `Service Name` as a default prefix.
func (dhs *HTTPService_) CustomPrefix(prefix string) {
	dhs.pathPrefix = prefix
}

func (dhs *HTTPService_) routes() map[string][]EndPoint {
	return dhs.routeMap
}

func (dhs *HTTPService_) middleware() *MiddlewareStack {
	return dhs.middlewareStack
}

// Use the specified Middleware for the `HTTPService`.
// The specified middleware will be inherited by any calls
// that are made on the HTTPService.
func (dhs *HTTPService_) Use(mw ...MiddlewareFunc) {
	dhs.middlewareStack.Use(mw...)
}

// Route maps a HTTP method request to the path and the specified handler.
func (dhs *HTTPService_) Route(path, httpMethod string, handler HandlerFunction) {
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

func NewHTTPService(name string) *HTTPService_ {
	ds := newService(name)
	routeMap := make(map[string][]EndPoint)
	var mwf []MiddlewareFunc
	mws := MiddlewareStack{
		stack: mwf,
	}
	dhs := HTTPService_{
		Service_:        &ds,
		routeMap:        routeMap,
		middlewareStack: &mws,
	}
	return &dhs
}
