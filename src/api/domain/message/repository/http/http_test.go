package http_test

import (
	"go-rest-chat/src/api/domain/message/repository/http"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewRepository_Success(t *testing.T) {
	// Given
	assert := assert.New(t)

	// When
	mockContainer, _ := dependencies.NewMockContainer()
	repository := http.NewRepository(mockContainer)

	// Then
	assert.NotNil(repository)
}
