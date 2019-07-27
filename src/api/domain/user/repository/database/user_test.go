package database_test

import (
	"errors"
	"go-rest-chat/src/api/domain/user/entities"
	"go-rest-chat/src/api/domain/user/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUser_Success(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer container.Database().Close()

	// When
	row := []map[string]interface{}{
		{"id": uint(2), "username": "test", "password": "test"},
	}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(row)

	// Then
	user, err := repository.GetUser("test")
	assert.Nil(err)
	assert.Equal("test", user.Username)
}

func Test_GetUser_UserNotFound_Fail(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer container.Database().Close()

	// When
	row := []map[string]interface{}{}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(row)

	// Then
	_, err := repository.GetUser("test")
	assert.NotNil(err)
	assert.EqualError(err, "User not found")
}

func Test_CreateUser_Success(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer container.Database().Close()

	// When
	user := entities.User{
		Username: "test",
		Password: "test",
	}

	container.Catcher().Reset().NewMock().WithQuery(`INSERT INTO "users"`)

	// Then
	userID, err := repository.CreateUser(user)
	assert.Nil(err)
	assert.Equal(uint(5577006791947779410), userID)
}

func Test_CreateUser_Fail(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer container.Database().Close()

	// When
	user := entities.User{
		Username: "test",
		Password: "test",
	}

	container.Catcher().Reset().NewMock().WithQuery(`INSERT`).WithError(errors.New("forced for test"))

	// Then
	_, err := repository.CreateUser(user)
	assert.NotNil(err)
	assert.EqualError(err, "forced for test")
}

func getDependenciesMock(t *testing.T) (*assert.Assertions, *dependencies.Container, repository.UserRepository) {
	assert := assert.New(t)
	mockContainer, err := dependencies.NewMockContainer()
	assert.Nil(err)
	repository := repository.NewUserRepository(mockContainer)
	return assert, mockContainer, repository
}
