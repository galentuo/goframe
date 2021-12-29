package goframe

import (
	"fmt"
	"sync"
)

type service struct {
	name string
	env  *sync.Map
}

func (ds *service) Name() string {
	return ds.name
}

func (ds *service) SetInCtx(key string, value interface{}) {
	ds.env.Store(key, value)
}

func (ds *service) getCtxData() *sync.Map {
	return ds.env
}

func newService(name string) service {
	return service{
		name: name,
	}
}

type httpService struct {
	*service
	pathPrefix      string
	routeMap        map[string][]EndPoint
	middlewareStack *MiddlewareStack
	children        []*httpService
}

func (dhs *httpService) prefix() string {
	if dhs.pathPrefix == "" {
		return fmt.Sprintf("/%s", dhs.name)
	}
	return dhs.pathPrefix
}

// CustomPrefix replaces the default path prefix by the
// custom one passed in. The routes on the service
// would have the `Service Name` as a default prefix.
func (dhs *httpService) CustomPrefix(prefix string) {
	dhs.pathPrefix = prefix
}

func (dhs *httpService) routes() map[string][]EndPoint {
	return dhs.routeMap
}

func (dhs *httpService) middleware() *MiddlewareStack {
	return dhs.middlewareStack
}

// Use the specified Middleware for the `HTTPService`.
// The specified middleware will be inherited by any calls
// that are made on the HTTPService.
func (dhs *httpService) Use(mw ...MiddlewareFunc) {
	dhs.middlewareStack.Use(mw...)
}

// Route maps a HTTP method request to the path and the specified handler.
func (dhs *httpService) Route(path, httpMethod string, handler HandlerFunction) {
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

func NewHTTPService(name string) *httpService {
	ds := newService(name)
	routeMap := make(map[string][]EndPoint)
	var mwf []MiddlewareFunc
	mws := MiddlewareStack{
		stack: mwf,
	}
	dhs := httpService{
		service:         &ds,
		routeMap:        routeMap,
		middlewareStack: &mws,
	}
	return &dhs
}

// Group creates a new `HTTPService` that inherits from it's parent HTTPService.
// This is useful for creating groups of end-points that need to share
// common functionality, like middleware.
/*
	a := NewHTTPService("a")
	g := a.NewGroup()
	g.Use(AuthorizeAPIMiddleware)
*/
func (dhs *httpService) NewGroup() *httpService {
	ms := *dhs.middlewareStack
	newHttpService := httpService{
		service:         dhs.service,
		pathPrefix:      dhs.pathPrefix,
		routeMap:        make(map[string][]EndPoint),
		middlewareStack: &ms,
	}
	children := []*httpService{}
	if dhs.children != nil {
		children = dhs.children
	}
	children = append(children, &newHttpService)
	dhs.children = children
	return &newHttpService
}

func (dhs *httpService) getChildren() []*httpService {
	return dhs.children
}
