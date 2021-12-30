package goframe

// MiddlewareFunc defines the interface for a piece of goframe
// Middleware.
/*
	func DoSomething(next HandlerFunction) HandlerFunction {
		return func(c ServerContext) error {
			// do something before calling the next handler
			err := next(c)
			// do something after call the handler
			return err
		}
	}
*/
type MiddlewareFunc func(HandlerFunction) HandlerFunction

// Use the specified Middleware for the `HTTPService`.
// The specified middleware will be inherited by any calls
// that are made on the HTTPService.
func (ms *MiddlewareStack) Use(mw ...MiddlewareFunc) {
	ms.stack = append(ms.stack, mw...)
}

// MiddlewareStack manages the middleware stack for an app/Group.
type MiddlewareStack struct {
	stack []MiddlewareFunc
}

func (ms *MiddlewareStack) handler(h HandlerFunction) HandlerFunction {
	if len(ms.stack) > 0 {
		mh := func(_ HandlerFunction) HandlerFunction {
			return h
		}

		tstack := []MiddlewareFunc{mh}
		tstack = append(tstack, ms.stack...)
		for _, mw := range tstack {
			h = mw(h)
		}
		return h
	}
	return h
}
