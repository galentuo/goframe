package goframe

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/galentuo/goframe/logger"
	"github.com/google/uuid"
)

func APIHandler(hf HandlerFunction, api HTTPService, path, method string, ll logger.LogLevel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl := &responseLogger{w: w, status: http.StatusOK}
		w = httpsnoop.Wrap(w, httpsnoop.Hooks{
			Write: func(httpsnoop.WriteFunc) httpsnoop.WriteFunc {
				return rl.Write
			},
			WriteHeader: func(httpsnoop.WriteHeaderFunc) httpsnoop.WriteHeaderFunc {
				return rl.WriteHeader
			},
		})
		ctx := NewServerContext(r.Context(), ll, w, r)

		t := time.Now()

		defer func() {
			ctx.Logger().WithFields(map[string]interface{}{
				"status": rl.status,
			}).Info(fmt.Sprintf("[%s] %s%s took %d ms", method, api.Prefix(), path, time.Since(t).Milliseconds()))
		}()
		log := ctx.Logger()
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		ctx.Set("service", api.Name())
		ctx.Set("req_id", reqID)
		log.SetField("req_id", reqID).
			SetField("service", api.Name()).
			SetField("path", r.URL.Path).
			SetField("method", r.Method)
		err := api.Middleware().handler(hf)(ctx)

		if err != nil {
			ctx.Logger().Error(err.Error())
			ctx.Response().ErrorJSON(errors.New("something went wrong"))
		}
	})
}

type responseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

func (l *responseLogger) Write(b []byte) (int, error) {
	size, err := l.w.Write(b)
	l.size += size
	return size, err
}

func (l *responseLogger) WriteHeader(s int) {
	l.w.WriteHeader(s)
	l.status = s
}
