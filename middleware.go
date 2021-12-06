package goframe

// MiddlewareFunc defines the interface for a piece of Buffalo
// Middleware.
/*
	func DoSomething(next HandlerFunction) HandlerFunction {
		return func(c Context) error {
			// do something before calling the next handler
			err := next(c)
			// do something after call the handler
			return err
		}
	}
*/
type MiddlewareFunc func(HandlerFunction) HandlerFunction

const funcKeyDelimeter = ":"

// Use the specified Middleware for the App.
// When defined on a `HTTPService` the specified middleware will be
// inherited by any `Group` calls that are made on that on
// the HTTPService.
func (ms *MiddlewareStack) Use(mw ...MiddlewareFunc) {
	ms.stack = append(ms.stack, mw...)
}

// MiddlewareStack manages the middleware stack for an App/Group.
type MiddlewareStack struct {
	stack []MiddlewareFunc
}

func (ms *MiddlewareStack) handler(h HandlerFunction) HandlerFunction {
	if len(ms.stack) > 0 {
		mh := func(_ HandlerFunction) HandlerFunction {
			return h
		}

		tstack := []MiddlewareFunc{mh}
		for _, mw := range tstack {
			h = mw(h)
		}
	}
	return h
}