package activity

import (
	"github.com/kataras/iris"
)

type ControllerActivity struct {
}

var Default = ControllerActivity{}

func (controller ControllerActivity) Register(application *iris.Application, path string) {

	activity := application.Party(path, func(context iris.Context) {
		context.Next()
	})
	activity.Get("/activities")
	activity.Get("/activity/{activity_id:int}")
	activity.Post("/activity")
	activity.Patch("/activity")
	activity.Post("/activity/products")
}
