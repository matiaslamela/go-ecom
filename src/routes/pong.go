package routes

import (
	"github.com/matiaslamela/go-ecom/src/controllers"

	"github.com/gin-gonic/gin"
)

func PongRoutes(router *gin.RouterGroup) { //controlador?
	api := router.Group("/pong")
	api.GET("/", controllers.Pong)
}
