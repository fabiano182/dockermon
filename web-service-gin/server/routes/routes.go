package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/fabiano182/dockermon/web-service-gin/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		containers := main.Group("containers")
		{
			containers.GET("/", controllers.HelloFromDockermon)
			containers.GET("/list", controllers.ListContainers)
			//containers.GET("/list/all", controllers.ListAllContainers)
		}
	}
	return router
}
