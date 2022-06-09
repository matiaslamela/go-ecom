package routes

import (
	"backgo/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) { //controlador?
	api := router.Group("/users")
	api.GET("/", controllers.GetUsers)
	api.POST("/", controllers.CreateUser)
}
