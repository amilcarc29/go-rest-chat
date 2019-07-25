package http

import (
	"go-rest-chat/src/api/domain/message/usecases"
	"go-rest-chat/src/api/infraestructure/dependencies"
)

// MessageHandler defines a message handler struct
type MessageHandler struct {
	usecases *usecases.UseCases
}

// NewMessageHandler returns a new message handler
func NewMessageHandler(container *dependencies.Container) *MessageHandler {
	usecases := usecases.NewUseCases(container)
	return &MessageHandler{
		usecases: usecases,
	}
}
