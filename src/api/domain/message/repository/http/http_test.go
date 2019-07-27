package http_test

import (
	"go-rest-chat/src/api/domain/message/repository/http"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewRepository_Success(t *testing.T) {
	assert := assert.New(t)
	mockContainer, _ := dependencies.NewMockContainer()
	repository := http.NewRepository(mockContainer)

	assert.NotNil(repository)
}
