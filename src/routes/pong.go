package routes

import (
	"backgo/src/controllers"

	"github.com/gin-gonic/gin"
)

func PongRoutes(router *gin.RouterGroup) { //controlador?
	api := router.Group("/pong")
	api.GET("/", controllers.Pong)
}
