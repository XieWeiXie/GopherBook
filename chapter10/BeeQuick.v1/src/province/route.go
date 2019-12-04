package province

import "github.com/kataras/iris"

type ControllerProvince struct {
}

var Default = ControllerProvince{}

func (controller ControllerProvince) RegisterWithOut(application *iris.Application, path string) {

	middleWare := func(ctx iris.Context) {
		ctx.Next()
	}

	province := application.Party(path, middleWare)
	province.Get("/provinces", getProvinceHandler)
	province.Get("/province", getOneHandler)

}
