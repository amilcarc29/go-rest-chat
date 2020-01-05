package repository_test

import (
	"go-rest-chat/src/api/domain/user/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserRepository_Success(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	mockContainer, _ := dependencies.NewMockContainer()
	repository := repository.NewUserRepository(mockContainer)

	// Then
	assert.NotNil(repository)
}
