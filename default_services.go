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
	endpoints  map[string][]EndPoint
	middleware *MiddlewareStack
}

func (dhs DefaultHTTPService) Prefix() string {
	return fmt.Sprintf("/%s%s", dhs.name, dhs.prefix)
}
func (dhs *DefaultHTTPService) Endpoints() map[string][]EndPoint {
	return dhs.endpoints
}
func (dhs *DefaultHTTPService) Middleware() *MiddlewareStack {
	return dhs.middleware
}

func NewHTTPService(name string) DefaultHTTPService {
	ds := newService(name)
	endpoints := make(map[string][]EndPoint)
	var mwf []MiddlewareFunc
	mws := MiddlewareStack{
		stack: mwf,
	}
	dhs := DefaultHTTPService{
		DefaultService: &ds,
		endpoints:      endpoints,
		middleware:     &mws,
	}
	return dhs
}
