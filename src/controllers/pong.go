package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) { //controlador?
	c.JSON(http.StatusOK, gin.H{
		"data": "pong",
	})
}
