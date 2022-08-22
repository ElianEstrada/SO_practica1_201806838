package main

import (
	"server/routes"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	app.Use(cors.Middleware(cors.Config{
		Origins: "*",
		Methods: "GET, PUT, POST, DELETE",
		//RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     false,
		ValidateHeaders: false,
	}))

	//app.GET("/vehicle", controllers.GetVehicles)

	router := app.Group("")
	routes.Routes(router)

	app.Run(":4000")
}
