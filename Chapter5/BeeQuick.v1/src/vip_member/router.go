package vip_member

import "github.com/kataras/iris"

type ControllerVipMember struct {
}

func (controller ControllerVipMember) Register(app *iris.Application, path string) {
	middleware := func(ctx iris.Context) {
		ctx.Next()
	}

	vipMember := app.Party(path, middleware)
	vipMember.Get("/vip_members", getVipMemberHandle)
	vipMember.Patch("/vip_member/{id:uint}", patchVipMemberHandle)
}
