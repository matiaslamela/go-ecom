package routes

import (
	"github.com/gin-gonic/gin"
)

func Server(app *gin.Engine) { //controlador?
	api := app.Group("/api/v1")
	PongRoutes(api)
	UserRoutes(api)
}
