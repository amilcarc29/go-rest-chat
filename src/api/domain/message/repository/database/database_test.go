package database_test

import (
	"go-rest-chat/src/api/domain/message/repository/database"
	"testing"
)

func Test_DatabaseRepository_Success(t *testing.T) {
	// Given
	assert, containerMocked, _ := getDependenciesMock(t)

	// When
	repository := database.NewRepository(containerMocked)

	// Then
	assert.NotNil(repository)
}
