package main

import (
	models "backgo/src/models"
	routes "backgo/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	models.Connect()
	routes.Server(app)
	app.Run(":3001")
}
