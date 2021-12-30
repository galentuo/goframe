package goframe

import (
	"fmt"
	"strings"
)

type internalError struct {
	httpCode int
	errCode  string
	message  string
}

// Error returns the stringified error
func (e internalError) Error() string {
	return e.errCode + ": " + e.message
}

// HttpCode returns the http code associated with the Error
func (e internalError) HttpCode() int {
	return e.httpCode
}

// ErrCode returns the errCode of the error
func (e internalError) ErrCode() string {
	return e.errCode
}

// Message returns the message of the error
func (e internalError) Message() string {
	return e.message
}

// CustomError returns an instance of error e with custom message
func (e internalError) MoreDetailed(additionalDetails ...string) *internalError {
	e_ := e
	e_.message = fmt.Sprintf("%s; %s", e.message, strings.Join(additionalDetails, ","))
	return &e_
}

// NewInternalError is used to define a new internal error
func NewInternalError(httpCode int, errCode string, errMessage string) *internalError {
	return &internalError{
		httpCode: httpCode,
		errCode:  errCode,
		message:  errMessage,
	}
}
