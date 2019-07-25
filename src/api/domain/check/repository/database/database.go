package database

import (
	"go-rest-chat/src/api/infraestructure/dependencies"

	"github.com/jinzhu/gorm"
)

// CheckDatabaseRepository defines a repository struct
type CheckDatabaseRepository struct {
	database *gorm.DB
}

// NewRepository returns a new Check repository
func NewRepository(container *dependencies.Container) *CheckDatabaseRepository {
	return &CheckDatabaseRepository{
		database: container.Database(),
	}
}
