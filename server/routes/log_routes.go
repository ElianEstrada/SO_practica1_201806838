package routes

import "github.com/gin-gonic/gin"

func logRoutes(superRouter *gin.RouterGroup) {
	logRouter := superRouter.Group("log")
	{
		logRouter.GET("/", nil)
	}
}
