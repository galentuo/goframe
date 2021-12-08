package goframe

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/galentuo/goframe/logger"
)

var (
	cl *logger.CoreLogger
)

func init() {
	cl = logger.NewCoreLogger()
}

type App struct {
	ll     logger.LogLevel
	name   string
	config configReader
	mux    Router
}

func (a *App) Name() string              { return a.name }
func (a *App) LogLevel() logger.LogLevel { return a.ll }

func (a *App) CustomCoreLogger(cl_ *logger.CoreLogger) {
	cl = cl_
}

// Config() returns the config reader.
// configs for the app are to be kept inside configs/ dir in app root dir.
// config values can be fetched by keys eg. "server.host"
// In production configs can be stored as env var
// eg. for app name "simple" -> simple_server_host
func (a *App) Config() configReader { return a.config }

// App is where it all happens!
//
// strictSlash defines the trailing slash behavior for new routes.
// When true, if the route path is "/path/", accessing "/path" will perform a redirect
// to the former and vice versa. In other words, your application will always
// see the path as specified in the route.
//
// customLogger can be null
func NewApp(name string, strictSlash bool) *App {
	a := App{
		name:   name,
		config: NewConfigReader(name, "./configs/", name, "_"),
		mux:    NewRouter(strictSlash),
	}

	ll := logger.LogLevelFromStr(a.config.GetString("log.level"))
	a.ll = ll
	return &a
}

func (app *App) Register(_svc Service) {
	if _svc.loglevel() == "" {
		_svc.SetLogLevel(app.ll)
	}
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
		cl.Fatal(fmt.Sprintf("Unknown service type for service %#v", svc))
	}

	if api != nil {
		for path, routes := range api.routes() {
			for _, endpoint := range routes {
				app.mux.Handle(endpoint.Method(), api.prefix()+path, APIHandler(endpoint.Handler(), api, path, endpoint.Method(), app.LogLevel()))
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

	cl.Info(fmt.Sprintf("Starting app on %s", srv.Addr))
	err := srv.ListenAndServe()
	if err != nil {
		cl.Fatal(err.Error())
	}
	return nil
}
