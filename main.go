package main

import (
	"github.com/gin-gonic/gin"
	enviroment "github.com/matiaslamela/go-ecom/src/env"
	models "github.com/matiaslamela/go-ecom/src/models"
	routes "github.com/matiaslamela/go-ecom/src/routes"
)

func main() {
	enviroment.GetEnvVars()
	app := gin.Default()
	models.Connect()
	routes.Server(app)
	app.Run(":" + enviroment.PORT)
}
