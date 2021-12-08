package main

import (
	"time"

	"github.com/galento/goframe/examples/simple/service"
	"github.com/galentuo/goframe"
)

func main() {
	app := goframe.NewApp("simple", true)

	app.Register(&service.UserService)
	app.Start(app.Config().GetString("server.host"), app.Config().GetInt("server.port"), 15*time.Second, 15*time.Second)
}
