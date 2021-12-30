package goframe

import (
	"sync"
)

// Service is the basic unit of a goframe app.
// It forms the base of any services like
// * HTTPService for a http server
// * BackgroundService for running workers
type Service interface {
	// Name is the name of the service;
	// this would be the prefix in case of HTTPService
	Name() string
	// SetInCtx sets required data into the service context;
	// It can be used to pass env elements like connections, configs, etc.
	SetInCtx(key string, value interface{})
	getCtxData() *sync.Map
}

// BackgroundService is a goframe Service used
// for running background workers.
type BackgroundService interface {
	Service
	Run() error
}

// HTTPService is a goframe Service used
// for running a http server.
type HTTPService interface {
	Service
	prefix() string
	// CustomPrefix replaces the default path prefix by the
	// custom one passed in. The routes on the service
	// would have the `Service Name` as a default prefix.
	CustomPrefix(string)
	routes() map[string][]endPoint
	Route(path, httpMethod string, handler Handler)
	middleware() *MiddlewareStack
	// Use the specified Middleware for the `HTTPService`.
	// The specified middleware will be inherited by any calls
	// that are made on the HTTPService.
	Use(mw ...MiddlewareFunc)
	// Group creates a new `HTTPService` that inherits from it's parent `HTTPService`.
	// This is useful for creating groups of end-points that need to share
	// common functionality, like middleware.
	/*
		g := a.Group()
		g.Use(AuthorizeAPIMiddleware)
	*/
	NewGroup() *httpService
	getChildren() []*httpService
}

// Handler is the basis for a HTTPService Endpoint. A Handler
// will be given a ServerContext interface that represents the
// given request/response. It is the responsibility of the
// Handler to handle the request/response correctly. This
// could mean rendering a template, JSON, etc... or it could
// mean returning an error.
/*
	func (c ServerContext) error {
		return c.Response().GenericJSON("Hello World!")
	}
*/
type Handler func(ServerContext) error

// endPoint is a type comprising of
// a route's http method and a Handler
/*
	endpoints["/users"] = []endPoint{
		{
			httpMethod: "GET"
			handler: UserListHandler
		},
	}
*/
type endPoint struct {
	httpMethod string
	handler    Handler
}

// Method returns the http method assciated with the endPoint
func (e endPoint) Method() string {
	return e.httpMethod
}

// Handler returns the http Handler assciated with the endPoint
func (e endPoint) Handler() Handler {
	return e.handler
}
