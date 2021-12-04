package goframe

import (
	"context"
	"net/http"

	"github.com/galentuo/goframe/logger"
)

type BasicContext interface {
	context.Context
	Logger() logger.Logger
	Set(string, interface{})
	Get(interface{}) interface{}
}

type ServerContext interface {
	BasicContext
	Request() *http.Request
	Response() ResponseWriter
}

type ResponseWriter interface {
	SuccessJSON(httpStatusCode int, data interface{}, message string)
	ErrorJSON(err error)
	GenericJSON(interface{})
}
