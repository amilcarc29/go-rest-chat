package database_test

import (
	"go-rest-chat/src/api/domain/user/repository/database"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DatabaseRepository_Success(t *testing.T) {
	assert := assert.New(t)
	mockContainer, err := dependencies.NewMockContainer()
	assert.Nil(err)
	repository := database.NewRepository(mockContainer)

	assert.NotNil(repository)
}
