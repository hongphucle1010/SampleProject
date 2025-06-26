package routes

import "github.com/kataras/iris/v12"

func PreflightHandler(app *iris.Application) {
	app.Options("/{path:path}", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.StatusCode(iris.StatusNoContent)
	})
}