package goframe

import "github.com/galentuo/goframe/logger"

var cl *logger.CoreLogger

func init() {
	cl = logger.NewCoreLogger()
}

type App struct {
	name   string
	config configReader
}

func (a App) Name() string          { return a.name }
func (a *App) Config() configReader { return a.config }

func NewApp(name string) *App {
	a := App{
		name:   name,
		config: NewConfigReader(name, "./configs/", name, "_"),
	}
	return &a
}
