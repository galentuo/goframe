package goframe

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type defaultResponseWriter struct {
	res http.ResponseWriter
}

type standardJSONResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
}

// Generic is used to write data([]byte) as the http response
func (drw *defaultResponseWriter) Generic(httpCode int, data []byte) error {
	drw.res.WriteHeader(httpCode)
	_, err := drw.res.Write(data)
	return err
}

// SuccessJSON is used to write data as response to a successful http request.
func (drw *defaultResponseWriter) SuccessJSON(httpCode int, data interface{}, message string) error {
	responseJson := standardJSONResponse{
		Success: true,
		Data:    data,
		Message: message,
	}
	response, err := json.Marshal(responseJson)
	if err != nil {
		cl.Error(fmt.Sprintf("[SuccessJSON] Error: %s; val: %+v", err.Error(), data))
		return err
	}

	drw.res.Header().Add("Content-Type", "application/json")
	drw.res.WriteHeader(httpCode)
	_, err = drw.res.Write(response)

	return err
}

// SuccessJSON is used to write err as response to a unsuccessful http request.
func (drw *defaultResponseWriter) ErrorJSON(err error) error {
	responseJson := standardJSONResponse{
		Success: false,
	}
	httpCode := int(500)

	gfErr := goframeError{}
	if ok := errors.As(err, &gfErr); ok {
		httpCode = gfErr.HttpCode()
		responseJson.ErrorCode = gfErr.ErrCode()
		responseJson.Message = gfErr.Message()
	} else {
		responseJson.ErrorCode = "dev"
		responseJson.Message = err.Error()
	}

	response, e := json.Marshal(responseJson)
	if e != nil {
		cl.Error(fmt.Sprintf("[ErrorJSON] Error: %s; val: %s", e.Error(), err.Error()))
		return e
	}
	drw.res.Header().Add("Content-Type", "application/json")
	drw.res.WriteHeader(httpCode)
	_, err = drw.res.Write(response)
	return err
}
