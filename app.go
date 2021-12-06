package goframe

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/galentuo/goframe/logger"
)

var cl *logger.CoreLogger

func init() {
	cl = logger.NewCoreLogger()
}

type App struct {
	name   string
	config configReader
	mux    Router
}

func (a App) Name() string          { return a.name }
func (a *App) Config() configReader { return a.config }

func NewApp(name string) *App {
	a := App{
		name:   name,
		config: NewConfigReader(name, "./configs/", name, "_"),
		mux:    NewRouter(),
	}
	return &a
}

func (app *App) Register(_svc Service) {
	var (
		api HTTPService
		bg  BackgroundService
	)
	switch svc := _svc.(type) {
	case HTTPService:
		api = svc
	case BackgroundService:
		bg = svc
	default:
		cl.Fatal(fmt.Sprintf("Unknown service type for service %s", svc.Name()))
	}

	if api != nil {
		for path, endpoints := range api.Endpoints() {
			for _, endpoint := range endpoints {
				app.mux.Handle(endpoint.Method(), api.Prefix()+path, APIHandler(endpoint.Handler(), api, path, endpoint.Method()))
			}
		}
	}

	if bg != nil {
		bg.Run()
	}
}

func (app App) Start(host string, port int, readTimeout, writeTimeout time.Duration) error {
	srv := &http.Server{
		Handler: app.mux,
		Addr:    host + ":" + strconv.Itoa(port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	log.Fatal(srv.ListenAndServe())
	return nil
}
