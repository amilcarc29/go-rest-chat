package database_test

import (
	"go-rest-chat/src/api/domain/message/repository/database"
	"testing"
)

func Test_DatabaseRepository_Success(t *testing.T) {
	assert, containerMocked, _ := getDependenciesMock(t)
	repository := database.NewRepository(containerMocked)

	assert.NotNil(repository)
}
