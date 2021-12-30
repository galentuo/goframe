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
	routes() map[string][]EndPoint
	Route(path, httpMethod string, handler HandlerFunction)
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

// HandlerFunction is the basis for a HTTPService Endpoint. A Handler
// will be given a ServerContext interface that represents the
// given request/response. It is the responsibility of the
// HandlerFunction to handle the request/response correctly. This
// could mean rendering a template, JSON, etc... or it could
// mean returning an error.
/*
	func (c ServerContext) error {
		return c.Response().GenericJSON("Hello World!")
	}
*/
type HandlerFunction func(ServerContext) error

// Endpoint is a type comprising of
// a route's http method and a HandlerFunction
/*
	endpoints["/users"] = []EndPoint{
		{
			httpMethod: "GET"
			handlerFunction: UserListHandlerFunction
		},
	}
*/
type EndPoint struct {
	httpMethod      string
	handlerFunction HandlerFunction
}

// Method returns the http method assciated with the EndPoint
func (e EndPoint) Method() string {
	return e.httpMethod
}

// Handler returns the http Handler assciated with the EndPoint
func (e EndPoint) Handler() HandlerFunction {
	return e.handlerFunction
}
