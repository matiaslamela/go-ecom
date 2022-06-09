package services

import (
	models "github.com/matiaslamela/go-ecom/src/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := models.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func insertUsers(input CreateUserInput, userChannel chan *models.User, errChannel chan error) {
	userBuffer := &models.User{Name: input.Name, Email: input.Email, Age: input.Age}
	result := models.DB.Create(&userBuffer)
	if result.Error != nil {
		userChannel <- nil
		errChannel <- result.Error
		return
	}
	userChannel <- userBuffer
	errChannel <- nil
	return
}

func CreateUser(input CreateUserInput) (*models.User, error) {
	userChannel := make(chan *models.User)
	errChannel := make(chan error)
	go insertUsers(input, userChannel, errChannel)
	createdUser := <-userChannel
	creationError := <-errChannel
	if creationError != nil {
		return nil, creationError
	}
	return createdUser, nil
}
