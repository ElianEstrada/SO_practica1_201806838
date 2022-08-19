package main

import (
	"github.com/gin-gonic/gin"
	"server/routes"
)

func main() {
	app := gin.New()

	router := app.Group("")
	routes.Routes(router)

	app.Run(":4000")
}
