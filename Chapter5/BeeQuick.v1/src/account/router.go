package account

import "github.com/kataras/iris"

var Default = ControllerForAccount{}

type ControllerForAccount struct {
}

func (controller ControllerForAccount) RegisterWithOut(app *iris.Application, path string) {
	middleware := func(context iris.Context) {
		context.Next()
	}

	account := app.Party(path, middleware)
	{
		account.Post("/register", registerHandle)
		account.Post("/sign", signHandle)

	}

}

func (controller ControllerForAccount) RegisterWith(app *iris.Application, path string) {
	middleware := func(context iris.Context) {
		context.Next()
	}

	account := app.Party(path, middleware)
	{
		account.Post("/logout", logoutHandle)
		account.Get("/account/{id:uint}", getAccountHandle)
	}
}
