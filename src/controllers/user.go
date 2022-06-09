package controllers

import (
	"backgo/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var input services.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, creationErr := services.CreateUser(input)
	if creationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": creationErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
