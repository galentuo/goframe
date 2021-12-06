package main

import (
	"time"

	"github.com/galentuo/goframe"
)

func main() {
	app := goframe.NewApp("hello", true)

	service1 := goframe.NewHTTPService("world")
	service1.Route("/", "GET", HelloHandler)

	service2 := goframe.NewHTTPService("service2")
	service2.Route("/get", "GET", Service2Handler)

	app.Register(service1)
	app.Register(service2)
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

func Service2Handler(ctx goframe.ServerContext) error {
	return ctx.Response().SuccessJSON(200, "{name:service2}", "up & running")
}