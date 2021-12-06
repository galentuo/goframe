package goframe

import "fmt"

// assert DefaultService satisfies Service interface
var _ Service = &DefaultService{}

type DefaultService struct {
	name string
}

func (ds DefaultService) Name() string {
	return ds.name
}

func newService(name string) DefaultService {
	return DefaultService{
		name: name,
	}
}

// assert DefaultService satisfies Service interface
var _ HTTPService = &DefaultHTTPService{}

type DefaultHTTPService struct {
	*DefaultService
	prefix     string
	routes     map[string][]EndPoint
	middleware *MiddlewareStack
}

func (dhs *DefaultHTTPService) Prefix() string {
	if dhs.prefix == "" {
		return fmt.Sprintf("/%s", dhs.name)
	}
	return dhs.prefix
}

func (dhs *DefaultHTTPService) CustomPrefix(prefix string) {
	dhs.prefix = prefix
}

func (dhs *DefaultHTTPService) Routes() map[string][]EndPoint {
	return dhs.routes
}

func (dhs *DefaultHTTPService) Middleware() *MiddlewareStack {
	return dhs.middleware
}

func (dhs *DefaultHTTPService) Route(path, httpMethod string, handler HandlerFunction) {
	endpoint := EndPoint{
		httpMethod:      httpMethod,
		handlerFunction: handler,
	}
	var (
		endpoints []EndPoint
		ok        bool
	)
	routes := dhs.Routes()
	if endpoints, ok = routes[path]; !ok {
		endpoints = []EndPoint{}
	}

	endpoints = append(endpoints, endpoint)
	dhs.routes[path] = endpoints
}

func NewHTTPService(name string) *DefaultHTTPService {
	ds := newService(name)
	routes := make(map[string][]EndPoint)
	var mwf []MiddlewareFunc
	mws := MiddlewareStack{
		stack: mwf,
	}
	dhs := DefaultHTTPService{
		DefaultService: &ds,
		routes:         routes,
		middleware:     &mws,
	}
	return &dhs
}
