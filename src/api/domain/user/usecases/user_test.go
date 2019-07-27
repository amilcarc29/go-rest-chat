package usecases_test

import (
	"errors"
	"go-rest-chat/src/api/domain/user/entities"
	"go-rest-chat/src/api/domain/user/usecases"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateUser_Success(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	row := []map[string]interface{}{}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(row)
	user := entities.User{
		Username: "test",
		Password: "test",
	}
	container.Catcher().NewMock().WithQuery(`INSERT INTO "users"`)

	// Then
	newUserID, err := usecases.CreateUser(user)
	assert.Nil(err)
	assert.Equal(uint(5577006791947779410), newUserID)
}

func Test_CreateUser_UsernameAlreadyExists_Fail(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	row := []map[string]interface{}{
		{"id": uint(2), "username": "test", "password": "test"},
	}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(row)
	user := entities.User{
		Username: "test",
		Password: "test",
	}

	// Then
	_, err := usecases.CreateUser(user)
	assert.EqualError(err, "Username already exists")
}

func Test_CreateUser_Fail(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	row := []map[string]interface{}{}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(row)
	user := entities.User{
		Username: "test",
		Password: "test",
	}
	container.Catcher().NewMock().WithQuery(`INSERT`).WithError(errors.New("forced for test"))

	// Then
	_, err := usecases.CreateUser(user)
	assert.NotNil(err)
	assert.EqualError(err, "forced for test")
}

func Test_LoginUser_Success(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	row := []map[string]interface{}{
		{"id": uint(2), "username": "test", "password": "$2a$05$hoDkIPIevK0kIXmBK8ilEuz6iKifccSFy9ndir9CU.z/1ZIqy4ZCu"},
	}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(row)
	user := entities.UserLogin{
		Username: "test",
		Password: "test",
	}

	// Then
	loginResponse, err := usecases.LoginUser(user)
	assert.Nil(err)
	assert.Equal(uint(2), loginResponse.ID)
	assert.NotZero(loginResponse.Token)
}

func Test_LoginUser_UserNotFound_Fail(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	row := []map[string]interface{}{}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(row)
	user := entities.UserLogin{
		Username: "notfoundtest",
		Password: "test",
	}

	// Then
	loginResponse, err := usecases.LoginUser(user)
	assert.NotNil(err)
	assert.EqualError(err, "invalid user")
	assert.Nil(loginResponse)
}

func Test_LoginUser_InvalidPassword_Fail(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	row := []map[string]interface{}{
		{"id": uint(2), "username": "test", "password": "$2a$05$hoDkIPIevK0kIXmBK8ilEuz6iKifccSFy9ndir9CU.z/1ZIqy4ZCu"},
	}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(row)
	user := entities.UserLogin{
		Username: "test",
		Password: "invalidpassword",
	}

	// Then
	loginResponse, err := usecases.LoginUser(user)
	assert.NotNil(err)
	assert.EqualError(err, "invalid user")
	assert.Nil(loginResponse)
}

func Test_AuthenticatedUser_Success(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	row := []map[string]interface{}{
		{"id": uint(2), "username": "test", "password": "$2a$05$hoDkIPIevK0kIXmBK8ilEuz6iKifccSFy9ndir9CU.z/1ZIqy4ZCu"},
	}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "users"`).WithReply(row)
	user := entities.UserLogin{
		Username: "test",
		Password: "test",
	}

	//Then
	loginResponse, _ := usecases.LoginUser(user)
	assert.NotNil(loginResponse)
	authenticated, err := usecases.AuthenticatedUser(loginResponse.Token)
	assert.Nil(err)
	assert.True(authenticated.Authenticated)
}

func Test_AuthenticatedUser_Fail(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	//Then
	authenticated, err := usecases.AuthenticatedUser("invalidToken")
	assert.NotNil(err)
	assert.False(authenticated.Authenticated)
}

func getDependenciesMock(t *testing.T) (*assert.Assertions, *dependencies.Container, *usecases.UseCases) {
	assert := assert.New(t)
	mockContainer, err := dependencies.NewMockContainer()
	assert.Nil(err)
	usecases := usecases.NewUseCases(mockContainer)
	return assert, mockContainer, usecases
}

func closeDependencies(container *dependencies.Container) {
	container.Database().Close()
}
