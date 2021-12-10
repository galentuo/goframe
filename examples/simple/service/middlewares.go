package service

import (
	"fmt"

	"github.com/galentuo/goframe"
)

var m1 = func(next goframe.HandlerFunction) goframe.HandlerFunction {
	return func(c goframe.ServerContext) error {
		// do something before calling the next handler
		fmt.Printf("\n\n >> Called middleware 1 << \n \n")
		err := next(c)
		// do something after calling the handler
		return err
	}
}

var m2 = func(next goframe.HandlerFunction) goframe.HandlerFunction {
	return func(c goframe.ServerContext) error {
		// do something before calling the next handler
		err := next(c)
		// do something after calling the handler
		fmt.Printf("\n\n >> Called middleware 2 << \n \n")
		return err
	}
}
