package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jfirme-sys/go-auth-docker.git/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.POST("/", controllers.CreateUser)
			users.POST("login", controllers.Login)
		}

	}

	return router
}
