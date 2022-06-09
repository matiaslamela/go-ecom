package services

import (
	models "backgo/src/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := models.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func CreateUser(input CreateUserInput) (*models.User, error) {
	user := &models.User{Name: input.Name, Email: input.Email, Age: input.Age}
	result := models.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
