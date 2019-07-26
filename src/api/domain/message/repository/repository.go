package repository

import (
	"go-rest-chat/src/api/domain/message/entities"
	database "go-rest-chat/src/api/domain/message/repository/database"
	"go-rest-chat/src/api/domain/message/repository/http"
	userEntities "go-rest-chat/src/api/domain/user/entities"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"time"
)

// MessageRepository defines an interface
type MessageRepository interface {
	GetMessages(sender, recipient, start, limit uint) ([]entities.Message, error)
	PutMessage(now time.Time, message entities.Message) (uint, time.Time, error)
	IsAuthenticated(token string) (userEntities.AuthenticatedResponse, error)
}

// Repository defines a repository struct
type Repository struct {
	*database.MessageDatabaseRepository
	*http.MessageHTTPRepository
}

// NewMessageRepository returns a new message repository
func NewMessageRepository(container *dependencies.Container) MessageRepository {
	return &Repository{
		database.NewRepository(container),
		http.NewRepository(container),
	}
}
