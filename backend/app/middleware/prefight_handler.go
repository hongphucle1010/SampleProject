package middleware

import "github.com/kataras/iris/v12"

func PreflightHandler(ctx iris.Context) {
	if ctx.Method() == iris.MethodOptions {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
	ctx.Next()
}
