package router_v1

import (
	"net/http"
	"time"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/order"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/product"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/tags"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/brand"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/unit"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/activity"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/shop"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/province"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/rule"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/exchange_coupons"

	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/vip_member"

	"github.com/kataras/iris"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/pkg/middleware"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/account"
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
	{

		account.Default.RegisterWithOut(app, "/v1")
		rule.Default.RegisterWithout(app, "/v1")
		province.Default.RegisterWithOut(app, "/v1")
		shop.Default.RegisterWithout(app, "/v1")
		activity.Default.Register(app, "/v1", false)
		unit.Default.Register(app, "/v1")
		brand.Default.Register(app, "/v1")
		tags.Default.Register(app, "/v1")
		product.Default.Register(app, "/v1")
	}

	app.Use(middleware.TokenForProject)

	{
		account.Default.RegisterWith(app, "/v1")
		vip_member.Default.Register(app, "/v1")
		exchange_coupons.Default.Register(app, "/v1")
		activity.Default.Register(app, "/v1", true)
		order.Default.Register(app, "/v1")
	}

	app.Logger().SetLevel("debug")
	return app
}

func notFound(app *iris.Application) {
	app.OnErrorCode(http.StatusNotFound, func(context iris.Context) {
		context.JSON(iris.Map{
			"code":   http.StatusNotFound,
			"detail": context.Request().URL.Path,
			"error":  "error found",
		})
	})
	return
}
