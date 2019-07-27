package repository_test

import (
	"go-rest-chat/src/api/domain/user/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserRepository_Success(t *testing.T) {
	assert := assert.New(t)
	mockContainer, _ := dependencies.NewMockContainer()
	repository := repository.NewUserRepository(mockContainer)

	assert.NotNil(repository)
}
