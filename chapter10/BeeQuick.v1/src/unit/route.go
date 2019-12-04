package unit

import "github.com/kataras/iris"

type ControllerUnit struct {
}

var Default = ControllerUnit{}

func (controller ControllerUnit) Register(application *iris.Application, path string) {
	unit := application.Party(path, func(context iris.Context) {
		context.Next()
	})
	unit.Get("/units", getUintHandle)
	unit.Post("/unit", createUintHandle)
	unit.Patch("/unit/{unit_id:int}", patchUintHandle)
}
