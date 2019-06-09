package router_v1

import (
	"net/http"
	"time"

	"GopherBook/Chapter5/BeeQuick.v1/pkg/middleware"
	"GopherBook/Chapter5/BeeQuick.v1/src/account"
	"github.com/kataras/iris"
)

var (
	VERSION = "v0.1.0"
)

func ApplyRouter() *iris.Application {
	app := iris.Default()

	notFound(app)

	app.Handle("GET", "/", func(context iris.Context) {
		_, _ = context.JSON(iris.Map{
			"data": time.Now().Format("2006-01-02 15:04:05"),
			"code": http.StatusOK,
		})
	})

	app.Handle("GET", "/heart", func(c iris.Context) {
		c.JSON(iris.Map{
			"data": time.Now().Format("2006-01-02 15:04:05"),
			"code": http.StatusOK,
		})
	})

	v1 := app.Party("/v1")
	v1.Get("/version", func(context iris.Context) {
		context.JSON(
			iris.Map{
				"code":    http.StatusOK,
				"version": VERSION,
			},
		)
		return
	})

	app.UseGlobal(middleware.LoggerForProject)
	var acc account.ControllerForAccount
	{

		acc.RegisterWithOut(app, "/v1")
	}

	app.Use(middleware.TokenForProject)
	{
		acc.RegisterWith(app, "/v1")
	}

	app.Logger().SetLevel("debug")
	return app
}

func notFound(app *iris.Application) {
	app.OnErrorCode(http.StatusNotFound, func(context iris.Context) {
		context.JSON(iris.Map{
			"code":   http.StatusNotFound,
			"detail": "error found",
		})
	})
	return
}
