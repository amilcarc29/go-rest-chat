package http

import (
	"go-rest-chat/src/api/domain/check/usecases"
	"go-rest-chat/src/api/infraestructure/dependencies"
)

type CheckHandler struct {
	usecases *usecases.UseCases
}

func NewCheckHandler(container *dependencies.Container) *CheckHandler {
	usecases := usecases.NewUseCases(container)
	return &CheckHandler{
		usecases: usecases,
	}
}
