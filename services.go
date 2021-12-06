package goframe

type Service interface {
	Name() string
}

type BackgroundService interface {
	Service
	Run() error
}

type HTTPService interface {
	Service
	Prefix() string
	Endpoints() map[string][]EndPoint
	Middleware() *MiddlewareStack
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
