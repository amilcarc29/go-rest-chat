package database

import (
	"go-rest-chat/src/api/infraestructure/dependencies"

	"github.com/jinzhu/gorm"
)

// UserDatabaseRepository defines a repository struct
type UserDatabaseRepository struct {
	database *gorm.DB
}

// NewRepository returns a new User repository
func NewRepository(container *dependencies.Container) *UserDatabaseRepository {
	return &UserDatabaseRepository{
		database: container.Database(),
	}
}
