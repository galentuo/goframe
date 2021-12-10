package goframe

// assert that defaultResponseWriter implementations is fulfilling its interface
var _ ResponseWriter = &defaultResponseWriter{}
