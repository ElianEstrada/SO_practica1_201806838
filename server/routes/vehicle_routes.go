package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func vehicleRoutes(superRoute *gin.RouterGroup) {
	vehicleRouter := superRoute.Group("/vehicle")
	{
		vehicleRouter.GET("/", controllers.GetVehicles)
		vehicleRouter.GET("/:id", controllers.GetVehicle)
		vehicleRouter.POST("/add", controllers.AddVehicle)
		vehicleRouter.PATCH("/update/:id", controllers.UpdateVehicle)
		vehicleRouter.DELETE("/delete/:id", controllers.DeleteVehicle)
	}
}
