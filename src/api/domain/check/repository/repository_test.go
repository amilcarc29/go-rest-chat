package repository_test

import (
	"github.com/stretchr/testify/assert"
	"go-rest-chat/src/api/domain/check/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"
)

func TestNewCheckRepository(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	container, _ := dependencies.NewMockContainer()
	repository := repository.NewCheckRepository(container)

	// Then
	ass.NotNil(repository)
}
