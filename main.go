package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	api := r.Group("/api")
	{
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/token", controllers.GenerateToken)
		api.GET("/ping", controllers.GetUser)
		// secured := api.Group("/secured").Use(controllers.Auth())
		// {
		// 	secured.GET("/ping", controllers.GetUser)
		// }
	}
	//return router
	r.Run(":8081")
}
