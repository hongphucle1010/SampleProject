package routes

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
)

func RegisterSwagger(app *iris.Application) {
	// Redirect "/" to Swagger UI
	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/swagger/index.html", iris.StatusTemporaryRedirect)
	})

	app.Get("/swagger/{any:path}", func(ctx iris.Context) {
		swagger.CustomWrapHandler(&swagger.Config{
			URL: "http://" + ctx.Host() + "/swagger/doc.json",
		}, swaggerFiles.Handler)(ctx)
	})
}
