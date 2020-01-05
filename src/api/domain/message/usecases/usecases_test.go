package usecases_test

import (
	"go-rest-chat/src/api/domain/message/usecases"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewUseCases(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	mockContainer, err := dependencies.NewMockContainer()
	assert.Nil(err)
	usecases := usecases.NewUseCases(mockContainer)

	// Then
	assert.NotNil(usecases)
}

func TestNewUseCasesMock(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	usecases := usecases.NewUseCasesMock(nil, nil)

	// Then
	assert.NotNil(usecases)
}
