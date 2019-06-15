package activity

import (
	"github.com/kataras/iris"
)

type ControllerActivity struct {
}

var Default = ControllerActivity{}

func (controller ControllerActivity) Register(application *iris.Application, path string, withToken bool) {

	activity := application.Party(path, func(context iris.Context) {
		context.Next()
	})

	if withToken {
		activity.Post("/activity", createOneActivityHandle)
		activity.Patch("/activity")
		activity.Post("/activity/products")
	} else {
		activity.Get("/activities")
		activity.Get("/activity/{activity_id:int}")
	}

}
