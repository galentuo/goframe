package goframe

import "github.com/galentuo/goframe/logger"

type Service interface {
	Name() string
	loglevel() string
	SetLogLevel(logger.LogLevel)
}

type BackgroundService interface {
	Service
	Run() error
}

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
