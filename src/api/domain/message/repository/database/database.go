package database

import (
	"go-rest-chat/src/api/infraestructure/dependencies"

	"github.com/jinzhu/gorm"
)

// MessageDatabaseRepository defines a repository struct
type MessageDatabaseRepository struct {
	database *gorm.DB
}

// NewRepository returns a new message repository
func NewRepository(container *dependencies.Container) *MessageDatabaseRepository {
	return &MessageDatabaseRepository{
		database: container.Database(),
	}
}
