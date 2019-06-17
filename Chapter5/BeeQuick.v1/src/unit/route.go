package unit

import "github.com/kataras/iris"

type ControllerUint struct {
}

var Default = ControllerUint{}

func (controller ControllerUint) Register(application *iris.Application, path string) {
	uint := application.Party(path, func(context iris.Context) {
		context.Next()
	})
	uint.Get("/units")
	uint.Post("/uint", createUintHandle)
	uint.Patch("/uint/{uint_id:int}")
}
