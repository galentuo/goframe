package goframe

import (
	"context"
	"net/http"
	"net/url"

	"github.com/galentuo/goframe/logger"
)

type Context interface {
	context.Context
	Logger() *logger.Logger
	Set(string, interface{})
	Get(interface{}) interface{}
}

type ServerContext interface {
	Context
	Request() *http.Request
	Response() ResponseWriter
	// Params returns all of the parameters for the request,
	// including both named params and query string parameters.
	Params() url.Values
	// Param returns a param, either named or query string,
	// based on the key.
	Param(key string) string
}

type ResponseWriter interface {
	SuccessJSON(httpStatusCode int, data interface{}, message string) error
	ErrorJSON(err error) error
	Generic(httpStatusCode int, data []byte) error
}
