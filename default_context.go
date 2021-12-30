package goframe

import (
	"context"
	"net/http"
	"net/url"
	"sync"

	"github.com/galentuo/goframe/logger"
)

// defaultContext is, as its name implies, a default
// implementation of the Context interface.
type defaultContext struct {
	context.Context
	data   *sync.Map
	logger *logger.Logger
}

// Logger returns the Logger for this context.
func (dbc defaultContext) Logger() *logger.Logger {
	return dbc.logger
}

// Set a value onto the Context.
func (dbc *defaultContext) Set(key string, value interface{}) {
	dbc.data.Store(key, value)
}

// Value that has previously been stored on the context.
func (dbc defaultContext) Get(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		if v, ok := dbc.data.Load(k); ok {
			return v
		}
	}
	return dbc.Context.Value(key)
}

// defaultServerContext is, as its name implies, a default
// implementation of the ServerContext interface.
type defaultServerContext struct {
	*defaultContext
	req    *http.Request
	res    http.ResponseWriter
	params url.Values
}

// Response returns goframe.ResponseWriter
func (dsc *defaultServerContext) Response() ResponseWriter {
	return &defaultResponseWriter{
		res: dsc.res,
	}
}

// Request returns *http.Request
func (dsc *defaultServerContext) Request() *http.Request {
	return dsc.req
}

// Params returns all of the parameters for the request,
// including both named params and query string parameters.
func (dsc *defaultServerContext) Params() url.Values {
	return dsc.params
}

// Param returns a param, either named or query string,
// based on the key.
func (dsc *defaultServerContext) Param(key string) string {
	return dsc.Params().Get(key)
}

// NewContext is used to get an instance of the default implementation
// of goframe.Context
func NewContext(ctx context.Context, ll logger.LogLevel) *defaultContext {
	dbc := defaultContext{
		data:    &sync.Map{},
		Context: ctx,
		logger:  logger.NewLogger(ll, cl, make(map[string]interface{})),
	}
	return &dbc
}

// NewServerContext is used to get an instance of the default implementation
// of goframe.ServerContext
func NewServerContext(ctx context.Context, ll logger.LogLevel, res http.ResponseWriter, req *http.Request) *defaultServerContext {
	llh := req.Header.Get("X-Request-LogLevel")
	if llh == "debug" {
		ll = logger.LogLevelDebug
	}
	dbc := NewContext(ctx, ll)
	dsc := defaultServerContext{
		defaultContext: dbc,
		res:            res,
		req:            req,
	}

	return &dsc
}
