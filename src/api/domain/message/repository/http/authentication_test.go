package http_test

import (
	"go-rest-chat/src/api/domain/message/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_IsAuthenticated_Success(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer httpmock.DeactivateAndReset()

	// When
	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(200, `{"authenticated": true, "id": "1"}`))

	authenticated, err := repository.IsAuthenticated("token")

	// Then
	assert.Nil(err)
	assert.True(authenticated.Authenticated)
}

func Test_IsAuthenticated_WithBearer_Success(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer httpmock.DeactivateAndReset()

	// When
	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(200, `{"authenticated": true, "id": "1"}`))

	authenticated, err := repository.IsAuthenticated("Bearer token")

	// Then
	assert.Nil(err)
	assert.True(authenticated.Authenticated)
}

func Test_IsAuthenticated_WithBearerButWithoutToken_Fail(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer httpmock.DeactivateAndReset()

	// When
	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(200, `{"authenticated": true, "id": "1"}`))

	authenticated, err := repository.IsAuthenticated("Bearer")

	// Then
	assert.NotNil(err)
	assert.EqualError(err, "invalid token")
	assert.False(authenticated.Authenticated)
}

func Test_IsAuthenticated_NotAuthenticated_Success(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer httpmock.DeactivateAndReset()

	// When
	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(401, `{"authenticated": false, "id": "0"}`))

	authenticated, err := repository.IsAuthenticated("Bearer token")

	// Then
	assert.NotNil(err)
	assert.EqualError(err, "not authenticated")
	assert.False(authenticated.Authenticated)
}

func getDependenciesMock(t *testing.T) (*assert.Assertions, *dependencies.Container, repository.MessageRepository) {
	assert := assert.New(t)
	mockContainer, err := dependencies.NewMockContainer()
	assert.Nil(err)
	repository := repository.NewMessageRepository(mockContainer)
	return assert, mockContainer, repository
}

func closeDependencies(container *dependencies.Container) {
	container.Database().Close()
	httpmock.DeactivateAndReset()
}
