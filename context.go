package goframe

import (
	"context"
	"net/http"

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
}

type ResponseWriter interface {
	SuccessJSON(httpStatusCode int, data interface{}, message string) error
	ErrorJSON(err error) error
	Generic(httpStatusCode int, data []byte) error
}
