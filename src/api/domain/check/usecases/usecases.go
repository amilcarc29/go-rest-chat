package usecases

import (
	"go-rest-chat/src/api/domain/check/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"
)

// UseCases defines a usecases struct
type UseCases struct {
	checkRepository repository.CheckRepository
}

// NewUseCases returns a new usecases
func NewUseCases(container *dependencies.Container) *UseCases {
	checkRepository := repository.NewCheckRepository(container)
	return &UseCases{
		checkRepository: checkRepository,
	}
}

// NewUseCasesMock returns a new usecases with repository mocked
func NewUseCasesMock(checkRepository repository.CheckRepository) *UseCases {
	return &UseCases{
		checkRepository: checkRepository,
	}
}
