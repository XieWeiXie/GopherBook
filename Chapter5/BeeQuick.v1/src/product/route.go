package product

import "github.com/kataras/iris"

type ControllerProduct struct {
}

var Default = ControllerProduct{}

func (c ControllerProduct) Register(application *iris.Application, path string) {
	product := application.Party(path, func(context iris.Context) {
		context.Next()
	})
	product.Get("/products")
	product.Get("/product/{product_id:int}")
	product.Post("/product")
	product.Patch("/product/{product_id:int}")
	product.Post("/product_multiply")
	product.Delete("/product/{product_id:int}")
}
