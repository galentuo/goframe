package goframe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// assert that DefaultResponseWriter implementations is fulfilling its interface
var _ ResponseWriter = &DefaultResponseWriter{}

type DefaultResponseWriter struct {
	res http.ResponseWriter
}

type standardJSONResponse struct {
	Status    bool        `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
}

func (drw DefaultResponseWriter) Generic(httpCode int, data []byte) error {
	drw.res.WriteHeader(httpCode)
	drw.res.Write(data)
	return nil
}

func (drw DefaultResponseWriter) SuccessJSON(httpCode int, data interface{}, message string) error {
	responseJson := standardJSONResponse{
		Status:  true,
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
	drw.res.Write(response)
	return nil
}

func (drw DefaultResponseWriter) ErrorJSON(err error) error {
	responseJson := standardJSONResponse{
		Status: false,
	}
	httpCode := int(500)

	if ie, ok := err.(*internalError); ok {
		httpCode = ie.HttpCode()
		responseJson.ErrorCode = ie.ErrCode()
		responseJson.Message = ie.Message()
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
	drw.res.Write(response)
	return nil
}
