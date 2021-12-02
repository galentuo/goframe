package main

import (
	"github.com/galentuo/goframe"
	"github.com/galentuo/goframe/examples/simple/pkg/config"
	"github.com/galentuo/goframe/examples/simple/service"
)

func main() {
	simpleConfig := config.Simple()
	server := goframe.NewApp()

	userService := service.NewUserService()
	userUpdater := service.NewUserUpdater()
	server.Register(userService)
	server.Register(userUpdater)
	server.Start(simpleConfig.Server.Host, simpleConfig.Server.Port)
}
