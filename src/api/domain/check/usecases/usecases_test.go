package usecases_test

import (
	"github.com/stretchr/testify/assert"
	"go-rest-chat/src/api/domain/check/usecases"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"
)

func TestNewUseCases(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	container, _ := dependencies.NewMockContainer()
	usecases := usecases.NewUseCases(container)

	// Then
	ass.NotNil(usecases)
}

func TestNewUseCasesMock(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When

	usecases := usecases.NewUseCasesMock(nil)

	// Then
	ass.NotNil(usecases)
}
