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
	shop.Get("/shops/{province_id:int}", getAllShopHandler)
	shop.Get("/shop/{shop_id:int}", getOneShopHandler)
	shop.Post("/shop", createOneShopHandler)
	shop.Patch("/shop/{shop_id:int}", patchOneShopHandler)
}
