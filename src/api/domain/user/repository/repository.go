package repository

import (
	"go-rest-chat/src/api/domain/user/entities"
	"go-rest-chat/src/api/domain/user/repository/database"
	"go-rest-chat/src/api/domain/user/repository/http"
	"go-rest-chat/src/api/infraestructure/dependencies"
)

// UserRepository defines an interface
type UserRepository interface {
	GetUser(username string) (entities.User, error)
	CreateUser(newUser entities.User) (uint, error)
}

// Repository defines a repository struct
type Repository struct {
	*database.UserDatabaseRepository
	*http.UserHTTPRepository
}

// NewUserRepository returns a new user repository
func NewUserRepository(container *dependencies.Container) UserRepository {
	return &Repository{
		database.NewRepository(container),
		http.NewRepository(container),
	}
}
