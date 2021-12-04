package goframe

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type DefaultResponseWriter struct {
	res http.ResponseWriter
}

type standardJSONResponse struct {
	Status    bool        `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
}

func (drw DefaultResponseWriter) GenericJSON(val interface{}) {
	response, err := json.Marshal(val)
	if err != nil {
		log.Panicln(fmt.Sprintf("Error: %s; val: %+v", err.Error(), val))
	}
	drw.res.Header().Add("Content-Type", "application/json")
	drw.res.Write(response)
}

func (drw DefaultResponseWriter) SuccessJSON(httpCode int, data interface{}, message string) {
	responseJson := standardJSONResponse{
		Status:  true,
		Data:    data,
		Message: message,
	}
	response, err := json.Marshal(responseJson)
	if err != nil {
		log.Panicln(fmt.Sprintf("Error: %s; val: %+v", err.Error(), data))
	}

	drw.res.Header().Add("Content-Type", "application/json")
	drw.res.WriteHeader(httpCode)
	drw.res.Write(response)
}

func (drw DefaultResponseWriter) ErrorJSON(err error) {
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
		log.Panicln(fmt.Sprintf("Error: %s; val: %+v", e.Error(), err.Error()))
	}
	drw.res.Header().Add("Content-Type", "application/json")
	drw.res.WriteHeader(httpCode)
	drw.res.Write(response)
}
