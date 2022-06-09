package main

import (
	models "github.com/matiaslamela/go-ecom/src/models"
	routes "github.com/matiaslamela/go-ecom/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	models.Connect()
	routes.Server(app)
	app.Run(":3001")
}
