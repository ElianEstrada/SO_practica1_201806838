package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func vehicleRoutes(superRoute *gin.RouterGroup) {
	vehicleRouter := superRoute.Group("/vehicle")
	{
		vehicleRouter.GET("/", controllers.GetVehicles)
	}
}
