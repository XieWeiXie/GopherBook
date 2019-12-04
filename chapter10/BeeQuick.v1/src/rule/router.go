package rule

import "github.com/kataras/iris"

type ControllerForRule struct {
}

var Default = ControllerForRule{}

func (c ControllerForRule) RegisterWithout(application *iris.Application, path string) {
	middleware := func(ctx iris.Context) {
		ctx.Next()
	}
	rule := application.Party(path, middleware)

	rule.Post("/rule", ruleCreateOneHandler)

	rule.Patch("/rule/{rule_id:int}", rulePatchOneHandler)

	rule.Get("/rule/{rule_id:int}", ruleGetOneHandler)

	rule.Get("/rules", ruleGetAllHandler)
}
