package services

import (
	"example.com/enterprise-grade-system/pkg/models"
	"example.com/enterprise-grade-system/pkg/repositories"
)

// GetUsers returns a list of users
:func GetUsers() ([]*models.User, error) {
	return repositories.GetUsers()
}

// GetUser returns a single user
:func GetUser(id string) (*models.User, error) {
	return repositories.GetUser(id)
}