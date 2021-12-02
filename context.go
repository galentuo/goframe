package goframe

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/galentuo/goframe/logger"
)

type Context interface {
	context.Context
	Logger() logger.Logger
	Set(string, interface{})
	Get(interface{}) interface{}
}

type ServerContext interface {
	Context
	Response() APIResponse
	Request() *http.Request
}

type APIContext struct {
	context.Context
	data   *sync.Map
	logger logger.Logger
	res    http.ResponseWriter
	req    *http.Request
}

type JSONStatus string

const (
	OK    JSONStatus = "ok"
	ERROR JSONStatus = "error"
)

type APIResponse interface {
	JSON(int, JSONStatus, interface{}, string)
	GenericJSON(interface{})
}

type JSONAPIResponse struct {
	res http.ResponseWriter
}

type standardJSONResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (js JSONAPIResponse) GenericJSON(val interface{}) {
	response, err := json.Marshal(val)
	if err != nil {

	}
	js.res.Header().Add("Content-Type", "application/json")
	js.res.Write(response)
}

func (js JSONAPIResponse) JSON(statusCode int, status JSONStatus, val interface{}, msg string) {
	responseJson := standardJSONResponse{
		Status:  string(status),
		Data:    val,
		Message: msg,
	}
	response, err := json.Marshal(responseJson)
	if err != nil {

	}
	js.res.Header().Add("Content-Type", "application/json")
	js.res.WriteHeader(statusCode)
	js.res.Write(response)
}

func (api APIContext) Logger() logger.Logger {
	return api.logger
}

func (api *APIContext) Set(key string, value interface{}) {
	api.data.Store(key, value)
}

// Value that has previously stored on the context.
func (api APIContext) Get(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		if v, ok := api.data.Load(k); ok {
			return v
		}
	}
	return api.Context.Value(key)
}

func (api APIContext) Response() APIResponse {
	return JSONAPIResponse{
		res: api.res,
	}
}

func (api APIContext) Request() *http.Request {
	return api.req
}

func NewAPIContext(ctx context.Context, cl *logger.CoreLogger, res http.ResponseWriter, req *http.Request) APIContext {
	llh := req.Header.Get("X-Request-LogLevel")
	ll := 1
	if llh == "debug" {
		ll = 0
	}
	return APIContext{
		data:    &sync.Map{},
		Context: ctx,
		res:     res,
		req:     req,
		logger:  logger.NewLogger(logger.LogLevel(ll), cl, make(map[string]interface{})),
	}
}
