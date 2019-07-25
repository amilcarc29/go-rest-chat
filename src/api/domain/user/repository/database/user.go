package database

import (
	"errors"
	"go-rest-chat/src/api/domain/user/entities"
)

const (
	// ErrUserNotFound defines
	ErrUserNotFound = "User not found"
)

// GetUser returns the user found
func (repository *UserDatabaseRepository) GetUser(username string) (entities.User, error) {
	var user entities.User
	err := repository.database.First(&user, entities.User{Username: username}).Error
	if user.Username == "" || err != nil {
		return user, errors.New(ErrUserNotFound)
	}
	return user, nil
}

// CreateUser inserts a new user
func (repository *UserDatabaseRepository) CreateUser(newUser entities.User) (uint, error) {
	if err := repository.database.Create(&newUser).Error; err != nil {
		return 0, err
	}
	return newUser.ID, nil
}
