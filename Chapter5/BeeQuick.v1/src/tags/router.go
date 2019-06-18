package tags

import "github.com/kataras/iris"

type ControllerTags struct {
}

var Default = ControllerTags{}

func (controller ControllerTags) Register(application *iris.Application, path string) {

	tags := application.Party(path, func(context iris.Context) {
		context.Next()
	})

	tags.Get("/tags", getTagsHandle)
	tags.Get("/tag/{tag_id:int}", getTagHandle)
	tags.Post("/tag", postTagHandle)
	tags.Post("/tag_multiply", postTagMultiplyHandle)
	tags.Patch("/tag/{tag_id:int}", patchTagHandle)
}
