package goframe

import (
	"context"
	"net/http"
	"net/url"

	"github.com/galentuo/goframe/logger"
)

// Context is goframe implementation of context.
// This is to be used across goframe
type Context interface {
	context.Context
	Logger() *logger.Logger
	Set(string, interface{})
	Get(interface{}) interface{}
}

// ServerContext is a goframe Context with more
// specific context related to a http server.
type ServerContext interface {
	Context
	// Request returns *http.Request
	Request() *http.Request
	// Response returns goframe.ResponseWriter
	Response() ResponseWriter
	// Params returns all of the parameters for the request,
	// including both named params and query string parameters.
	Params() url.Values
	// Param returns a param, either named or query string,
	// based on the key.
	Param(key string) string
}
