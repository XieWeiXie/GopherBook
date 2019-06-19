package product

import "github.com/kataras/iris"

type ControllerProduct struct {
}

var Default = ControllerProduct{}

func (c ControllerProduct) Register(application *iris.Application, path string) {
	product := application.Party(path, func(context iris.Context) {
		context.Next()
	})
	product.Get("/products", getAllProductHandle)
	product.Get("/product/{product_id:int}", getOneProductHandle)
	product.Post("/product", createProductHandle)
	product.Patch("/product/{product_id:int}", patchOneProductHandle)
	product.Post("/product_multiply", postMultiplyProductHandle)
	product.Delete("/product/{product_id:int}", deleteProductHandle)
}
