package services

import (
	"backend/config"
	"backend/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func CreateUser(user *models.User) error {
	result := config.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
