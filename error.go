package goframe

import (
	"fmt"
	"strings"
)

type goframeError struct {
	httpCode int
	errCode  string
	message  string
}

// Error returns the stringified error
func (e goframeError) Error() string {
	return e.errCode + ": " + e.message
}

// HttpCode returns the http code associated with the Error
func (e goframeError) HttpCode() int {
	return e.httpCode
}

// ErrCode returns the errCode of the error
func (e goframeError) ErrCode() string {
	return e.errCode
}

// Message returns the message of the error
func (e goframeError) Message() string {
	return e.message
}

// CustomError returns an instance of error e with custom message
func (e goframeError) MoreDetailed(additionalDetails ...string) *goframeError {
	e_ := e
	e_.message = fmt.Sprintf("%s; %s", e.message, strings.Join(additionalDetails, ","))
	return &e_
}

// NewError is used to define a new goframe error
func NewError(httpCode int, errCode string, errMessage string) *goframeError {
	return &goframeError{
		httpCode: httpCode,
		errCode:  errCode,
		message:  errMessage,
	}
}
