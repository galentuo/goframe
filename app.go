package goframe

import "github.com/galentuo/goframe/logger"

var cl *logger.CoreLogger

type App struct {
}

func NewApp() *App {
	cl = logger.NewCoreLogger()
	a := App{}
	return &a
}
