package brand

import "github.com/kataras/iris"

type ControllerBrand struct {
}

var Default = ControllerBrand{}

func (controller ControllerBrand) Register(application *iris.Application, path string) {
	brand := application.Party(path, func(context iris.Context) {
		context.Next()
	})

	brand.Get("/brands", getBrandsHandle)
	brand.Get("/brand/{brand_id:int}", getBrandHandle)
	brand.Post("/brand", createBrandHandle)
	brand.Patch("/brand/{brand_id:int}", patchBrandHandle)
	brand.Post("/brands_multiply", createBrandsHandle)

}
