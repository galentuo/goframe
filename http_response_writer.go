package goframe

// ResponseWriter interface is used by a goframe Handler
// to construct an HTTP response.
//
// A ResponseWriter may not be used after the Router.ServeHTTP method
// has returned.
type ResponseWriter interface {
	SuccessJSON(httpStatusCode int, data interface{}, message string) error
	ErrorJSON(err error) error
	Generic(httpStatusCode int, data []byte) error
}
