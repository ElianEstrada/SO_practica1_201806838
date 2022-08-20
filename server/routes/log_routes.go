package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func logRoutes(superRouter *gin.RouterGroup) {
	logRouter := superRouter.Group("log")
	{
		logRouter.POST("/add", controllers.AddLog)
	}
}
