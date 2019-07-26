package usecases

import (
	"go-rest-chat/src/api/domain/message/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"
)

// UseCases defines a usecases struct
type UseCases struct {
	messageRepository repository.MessageRepository
	clock             dependencies.Clock
}

// NewUseCases returns a new usecases
func NewUseCases(container *dependencies.Container) *UseCases {
	messageRepository := repository.NewMessageRepository(container)
	return &UseCases{
		messageRepository: messageRepository,
		clock:             container.Clock(),
	}
}
