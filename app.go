package goframe

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/galentuo/goframe/logger"
)

var cl *logger.CoreLogger

type App struct {
	mux Router
}

func NewApp() *App {
	cl = logger.NewCoreLogger()

	return &App{
		mux: NewRouter(),
	}
}

func (app *App) Register(_svc Service) {
	var (
		api RESTService
		bg  BackgroundService
	)
	switch svc := _svc.(type) {
	case RESTService:
		api = svc
	case BackgroundService:
		bg = svc
	default:
		cl.Fatal(fmt.Sprintf("Unknown servie type for service %s", svc.Name()))
	}

	if api != nil {
		for path, methodHandler := range api.Endpoints() {
			for method, handler := range methodHandler {
				app.mux.Handle(method, api.Prefix()+path, APIHandler(handler, api, path, method))
			}
		}
	}

	if bg != nil {
		bg.Run()
	}

}

func (app App) Start(host, port string) error {
	srv := &http.Server{
		Handler: app.mux,
		Addr:    fmt.Sprintf("%s:%s", host, port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	return nil
}
