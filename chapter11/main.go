package main

import (
	"GopherBook/chapter11/cmd"

	"github.com/kataras/iris/middleware/logger"

	"github.com/kataras/iris"
)

func main() {
	cmd.Execute()
	app := iris.New()
	app.Use(logger.New(logger.DefaultConfig()))
	app.Run(iris.Addr(":8080"))
}
