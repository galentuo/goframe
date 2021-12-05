package main

import (
	"fmt"

	"github.com/galentuo/goframe"
)

func main() {
	app := goframe.NewApp("simple")
	fmt.Println(app.Config().GetString("env"))
}
