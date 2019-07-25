package repository

import (
	"go-rest-chat/src/api/domain/check/repository/database"
	"go-rest-chat/src/api/infraestructure/dependencies"
)

// CheckRepository defines an interface
type CheckRepository interface {
	Check() (bool, error)
}

// Repository defines a repository struct
type Repository struct {
	*database.CheckDatabaseRepository
}

// NewCheckRepository returns a new Check repository
func NewCheckRepository(container *dependencies.Container) CheckRepository {
	return &Repository{
		database.NewRepository(container),
	}
}
