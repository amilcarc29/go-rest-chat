package http

import (
	"go-rest-chat/src/api/domain/user/usecases"
	"go-rest-chat/src/api/infraestructure/dependencies"
)

type UserHandler struct {
	usecases *usecases.UseCases
}

func NewUserHandler(container *dependencies.Container) *UserHandler {
	usecases := usecases.NewUseCases(container)
	return &UserHandler{
		usecases: usecases,
	}
}
