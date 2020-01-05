package usecases_test

import (
	"go-rest-chat/src/api/domain/user/usecases"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewUseCases_Success(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	mockContainer, err := dependencies.NewMockContainer()
	assert.Nil(err)
	usecases := usecases.NewUseCases(mockContainer)

	// Then
	assert.NotNil(usecases)
}

func Test_NewUseCasesMock_Success(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	usecases := usecases.NewUseCasesMock(nil)

	// Then
	assert.NotNil(usecases)
}
