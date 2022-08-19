package routes

import "github.com/gin-gonic/gin"

func Routes(groupRoute *gin.RouterGroup) {
	vehicleRoutes(groupRoute)
	logRoutes(groupRoute)
}
