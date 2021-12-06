package goframe

import (
	"context"
	"net/http"
	"sync"

	"github.com/galentuo/goframe/logger"
)

// assert that DefaultContext implementations are fulfilling their interfaces & context.Context
var _ context.Context = &DefaultServerContext{}
var _ context.Context = &DefaultBasicContext{}
var _ BasicContext = &DefaultBasicContext{}
var _ ServerContext = &DefaultServerContext{}

// DefaultBasicContext is, as its name implies, a default
// implementation of the BasicContext interface.
type DefaultBasicContext struct {
	context.Context
	data   *sync.Map
	logger *logger.Logger
}

// Logger returns the Logger for this context.
func (dbc DefaultBasicContext) Logger() *logger.Logger {
	return dbc.logger
}

// Set a value onto the Context.
func (dbc *DefaultBasicContext) Set(key string, value interface{}) {
	dbc.data.Store(key, value)
}

// Value that has previously been stored on the context.
func (dbc DefaultBasicContext) Get(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		if v, ok := dbc.data.Load(k); ok {
			return v
		}
	}
	return dbc.Context.Value(key)
}

// DefaultBasicContext is, as its name implies, a default
// implementation of the BasicContext interface.
type DefaultServerContext struct {
	*DefaultBasicContext
	req *http.Request
	res http.ResponseWriter
}

func (dsc DefaultServerContext) Response() ResponseWriter {
	return DefaultResponseWriter{
		res: dsc.res,
	}
}

func (dsc DefaultServerContext) Request() *http.Request {
	return dsc.req
}

func NewBasicContext(ctx context.Context, cl *logger.CoreLogger, ll logger.LogLevel) *DefaultBasicContext {
	dbc := DefaultBasicContext{
		data:    &sync.Map{},
		Context: ctx,
		logger:  logger.NewLogger(ll, cl, make(map[string]interface{})),
	}
	return &dbc
}

func NewServerContext(ctx context.Context, ll logger.LogLevel, res http.ResponseWriter, req *http.Request) DefaultServerContext {
	llh := req.Header.Get("X-Request-LogLevel")
	if llh == "debug" {
		ll = logger.LogLevelDebug
	}
	dbc := NewBasicContext(ctx, cl, ll)
	dsc := DefaultServerContext{
		DefaultBasicContext: dbc,
		res:                 res,
		req:                 req,
	}

	return dsc
}
