package database_test

import (
	"github.com/stretchr/testify/assert"
	"go-rest-chat/src/api/domain/check/repository/database"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"
)

func TestNewRepository(t *testing.T) {
	// Given
	ass := assert.New(t)

	container, _ := dependencies.NewMockContainer()

	// When
	repository := database.NewRepository(container)

	// Then
	ass.NotNil(repository)
}
