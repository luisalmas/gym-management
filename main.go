package main

import (
	"gym-management/controllers"

	_ "gym-management/docs" //import generated swagger doc files

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     Gym management API
// @version         1.0
// @description     A book management service API in Go using Gin framework.

// @contact.name   Luís Almas
// @contact.email  la_luisalmas@hotmail.com

// @host      localhost:8080
// @BasePath  /api
func main() {
    router := gin.Default()

	controllers.NewClassesController().SetupRoutes(router)

	//============= Swagger ==============================
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    router.Run("localhost:8080")
}