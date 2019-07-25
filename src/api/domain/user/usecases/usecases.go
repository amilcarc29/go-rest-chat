package usecases

import (
	"go-rest-chat/src/api/domain/user/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"
)

// UseCases defines a usecases struct
type UseCases struct {
	userRepository repository.UserRepository
}

// NewUseCases returns a new usecases
func NewUseCases(container *dependencies.Container) *UseCases {
	userRepository := repository.NewUserRepository(container)
	return &UseCases{
		userRepository: userRepository,
	}
}
