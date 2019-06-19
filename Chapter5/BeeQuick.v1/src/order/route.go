package order

import "github.com/kataras/iris"

type ControllerOrder struct {
}

var Default = ControllerOrder{}

func (c ControllerOrder) Register(application *iris.Application, path string) {
	order := application.Party(path, func(context iris.Context) {
		context.Next()
	})

	order.Get("/orders", getAllOrderHandle)
	order.Post("/order", postOrderHandle)
	order.Patch("/order/{order_id:int}", patchOrderHandle)
	order.Get("/order/{order_id:int}", getOneOrderHandle)
}
