package usecases_test

import (
	"go-rest-chat/src/api/domain/user/usecases"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewUseCases_Success(t *testing.T) {
	assert := assert.New(t)
	mockContainer, err := dependencies.NewMockContainer()
	assert.Nil(err)
	usecases := usecases.NewUseCases(mockContainer)

	assert.NotNil(usecases)
}
