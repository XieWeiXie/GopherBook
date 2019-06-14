package shop

import (
	"github.com/kataras/iris"
)

type ControllerShop struct {
}

var Default = ControllerShop{}

func (controller ControllerShop) RegisterWithout(application *iris.Application, path string) {
	middleware := func(ctx iris.Context) {
		ctx.Next()
	}
	shop := application.Party(path, middleware)
	shop.Get("/shops")
	shop.Get("/shop/{shop_id:int}")
	shop.Post("/shop")
	shop.Patch("/shop")
}
