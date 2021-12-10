package main

import (
	"time"

	"github.com/galentuo/goframe"
	"github.com/galentuo/goframe/examples/simple/service"
)

func main() {
	app := goframe.NewApp("simple", true, nil)

	app.Register(service.NewUserService().Service)
	app.Start(app.Config().GetString("server.host"), app.Config().GetInt("server.port"), 15*time.Second, 15*time.Second)
}
