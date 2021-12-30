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

type app struct {
	ll     logger.LogLevel
	name   string
	config Config
	mux    Router
}

func (a *app) Name() string              { return a.name }
func (a *app) LogLevel() logger.LogLevel { return a.ll }

func (a *app) CustomCoreLogger(clIn *logger.CoreLogger) {
	cl = clIn
}

// Config returns the config reader.
// Config values can be fetched by keys eg. "server.host".
// In production configs can be stored as env vars.
//
// Defaults:
// * The default config reader expects the config files
// 	 to be kept inside configs/ dir in app root dir.
// * The name of the app would be the default expected config file name.
// * While using env vars, `_` would be the default separator.
// * The env vars would have the app name as a default prefix.
// * eg. for app name "simple" -> simple_server_host
func (a *app) Config() Config { return a.config }

// NewApp returns an instance of app.
// app is where it all happens!
//
// strictSlash defines the trailing slash behavior for new routes.
// When true, if the route path is "/path/", accessing "/path" will perform a redirect
// to the former and vice versa. In other words, your application will always
// see the path as specified in the route.
//
// Config is a nullable field; if null it uses a default
// configReader = NewConfigReader(app.name, "./configs/", app.name, "_")
func NewApp(name string, strictSlash bool, cr Config) *app {
	if cr == nil {
		cr = NewConfig(name, "./configs/", name, "_")
	}
	a := app{
		name:   name,
		config: cr,
		mux:    NewRouter(strictSlash),
	}

	ll := logger.LogLevelFromStr(a.config.GetString("log.level"))
	a.ll = ll
	return &a
}

func (a *app) Register(_svc Service) {
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
				a.mux.Handle(endpoint.Method(), api.prefix()+path, apiHandler(endpoint.Handler(), api, path, endpoint.Method(), a.LogLevel(), api.getCtxData()))
			}
		}
		for _, each := range api.getChildren() {
			a.Register(each)
		}
	}

	if bg != nil {
		err := bg.Run()
		if err != nil {
			cl.Error(err.Error())
		}
	}
}

func (a *app) Start(host string, port int, readTimeout, writeTimeout time.Duration) error {
	srv := &http.Server{
		Handler: a.mux,
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
