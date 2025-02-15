package repository_test

import (
	"go-rest-chat/src/api/domain/message/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MessageRepository_Success(t *testing.T) {
	assert := assert.New(t)
	mockContainer, _ := dependencies.NewMockContainer()
	repository := repository.NewMessageRepository(mockContainer)

	assert.NotNil(repository)
}
