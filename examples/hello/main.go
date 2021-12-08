package main

import (
	"time"

	"github.com/galentuo/goframe"
)

func main() {
	app := goframe.NewApp("hello", true, nil)

	service1 := goframe.NewHTTPServer("")
	service1.Route("/", "GET", HelloHandler)

	app.Register(service1)
	app.Start(app.Config().GetString("server.host"), app.Config().GetInt("server.port"), 15*time.Second, 15*time.Second)
}

func HelloHandler(ctx goframe.ServerContext) error {
	msg := `
		<!DOCTYPE html>
		<html>
		<body>
		<h1 style="background-color:DodgerBlue;">Hello</h1>
		<h2 style="background-color:MediumSeaGreen;">World!</h2>
		</body>
		</html>
	`
	return ctx.Response().Generic(200, []byte(msg))
}
