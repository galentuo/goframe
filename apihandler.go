package goframe

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/galentuo/goframe/logger"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func APIHandler(hf HandlerFunction, api HTTPService, path, method string, ll logger.LogLevel, env *sync.Map) http.Handler {
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
		env.Range(func(key, value interface{}) bool {
			ctx.Set(fmt.Sprint(key), value)
			return true
		})
		t := time.Now()

		defer func() {
			ctx.Logger().WithFields(map[string]interface{}{
				"status": rl.status,
			}).Info(fmt.Sprintf("[%s] %s%s took %d ms", method, api.prefix(), path, time.Since(t).Milliseconds()))
		}()
		log := ctx.Logger()
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		w.Header().Add("X-Request-ID", reqID)

		// Parse URL Params
		params := url.Values{}
		vars := mux.Vars(r)
		for k, v := range vars {
			params.Add(k, v)
		}
		// Parse URL Query String Params
		// For POST, PUT, and PATCH requests, it also parse the request body as a form.
		// Request body parameters take precedence over URL query string values in params
		if err := r.ParseForm(); err == nil {
			for k, v := range r.Form {
				for _, vv := range v {
					params.Add(k, vv)
				}
			}
		}
		ctx.params = params

		ctx.Set("service", api.Name())
		ctx.Set("req_id", reqID)
		log.SetField("req_id", reqID).
			SetField("service", api.Name()).
			SetField("path", r.URL.Path).
			SetField("method", r.Method)

		err := api.middleware().handler(hf)(ctx)
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
