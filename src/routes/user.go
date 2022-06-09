package routes

import (
	"github.com/matiaslamela/go-ecom/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) { //controlador?
	api := router.Group("/users")
	api.GET("/", controllers.GetUsers)
	api.POST("/", controllers.CreateUser)
}
