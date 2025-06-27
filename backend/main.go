package main

import (
	"sample/app/config"
	"sample/app/middleware"
	"sample/app/routes"
	"sample/app/utils"
	_ "sample/docs"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

// @title Swagger UI Hệ thống quản lý sinh viên
// @version 1.0
// @description Tài liệu API của hệ thống quản lý sinh viên
// @termsOfService http://swagger.io/terms/

// @contact.name HCMUT Team tại FPT Software HCM
// @contact.email hongphucle1010@gmail.com

// @BasePath /api
func main() {
	// Initialize config
	config.InitConfig()

	// Initialize MongoDB
	utils.ConnectMongoDB(viper.GetString("database.uri"))

	// Initialize Iris
	app := iris.New()
	app.Use(middleware.Recover) // Recover from panics
	app.Use(middleware.Logger)  // Log requests
	app.Use(middleware.Cors())  // Handle CORS

	routes.Register(app) // Register routes

	// Start server
	app.Listen(":" + viper.GetString("server.port"))
}
