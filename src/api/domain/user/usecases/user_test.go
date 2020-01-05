package usecases_test

import (
	"errors"
	"go-rest-chat/src/api/domain/user/entities"
	userRepository "go-rest-chat/src/api/domain/user/repository"
	"go-rest-chat/src/api/domain/user/repository/database"
	"go-rest-chat/src/api/domain/user/repository/mocks"
	"go-rest-chat/src/api/domain/user/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateUser(t *testing.T) {
	ass := assert.New(t)

	tests := []struct {
		name       string
		repository func() userRepository.UserRepository
		newUser    entities.User
		asserts    func(id uint, err error)
	}{
		{
			name: "Success",
			repository: func() userRepository.UserRepository {
				repository := mocks.NewMock()

				repository.PatchGetUser("test", entities.User{}, errors.New(database.ErrUserNotFound))

				newUser := entities.User{
					Username: "test",
					Password: "test",
				}
				repository.PatchCreateUser(newUser, 1, nil)

				return repository
			},
			newUser: entities.User{
				Username: "test",
				Password: "test",
			},
			asserts: func(id uint, err error) {
				ass.Equal(uint(1), id)
				ass.Nil(err)
			},
		},
		{
			name: "Should Fail - Username already exists",
			repository: func() userRepository.UserRepository {
				repository := mocks.NewMock()
				user := entities.User{
					ID:       uint(2),
					Username: "test",
					Password: "test",
				}
				repository.PatchGetUser("test", user, nil)

				return repository
			},
			newUser: entities.User{
				Username: "test",
				Password: "test",
			},
			asserts: func(id uint, err error) {
				ass.Zero(id)
				ass.NotNil(err)
				ass.EqualError(err, "Username already exists")
			},
		},
		{
			name: "Should Fail - Error when getting user",
			repository: func() userRepository.UserRepository {
				repository := mocks.NewMock()
				user := entities.User{
					ID:       uint(2),
					Username: "test",
					Password: "test",
				}
				repository.PatchGetUser("test", user, errors.New("forced for test"))

				return repository
			},
			newUser: entities.User{
				Username: "test",
				Password: "test",
			},
			asserts: func(id uint, err error) {
				ass.Zero(id)
				ass.NotNil(err)
				ass.EqualError(err, "forced for test")
			},
		},
		{
			name: "Should Fail - Error inserting in DB",
			repository: func() userRepository.UserRepository {
				repository := mocks.NewMock()
				repository.PatchGetUser("test", entities.User{}, errors.New(database.ErrUserNotFound))

				newUser := entities.User{
					Username: "test",
					Password: "test",
				}
				repository.PatchCreateUser(newUser, 0, errors.New("forced for test"))
				return repository
			},
			newUser: entities.User{
				Username: "test",
				Password: "test",
			},
			asserts: func(id uint, err error) {
				ass.Zero(id)
				ass.NotNil(err)
				ass.EqualError(err, "forced for test")
			},
		},
	}

	for _, test := range tests {
		// Given
		repository := test.repository()

		// When
		usecases := usecases.NewUseCasesMock(repository)
		id, err := usecases.CreateUser(test.newUser)

		// Then
		test.asserts(id, err)
	}
}

func Test_LoginUser(t *testing.T) {
	ass := assert.New(t)

	tests := []struct {
		name       string
		repository func() userRepository.UserRepository
		loginUser  entities.UserLogin
		asserts    func(loginResponse *entities.LoginResponse, err error)
	}{
		{
			name: "Success",
			repository: func() userRepository.UserRepository {
				repository := mocks.NewMock()
				user := entities.User{
					ID:       uint(2),
					Username: "test",
					Password: "$2a$05$hoDkIPIevK0kIXmBK8ilEuz6iKifccSFy9ndir9CU.z/1ZIqy4ZCu",
				}
				repository.PatchGetUser("test", user, nil)

				return repository
			},
			loginUser: entities.UserLogin{
				Username: "test",
				Password: "test",
			},
			asserts: func(loginResponse *entities.LoginResponse, err error) {
				ass.NotNil(loginResponse)
				ass.Equal(uint(2), loginResponse.ID)
				ass.NotZero(loginResponse.Token)
				ass.Nil(err)
			},
		},
		{
			name: "Should Fail - Username not found",
			repository: func() userRepository.UserRepository {
				repository := mocks.NewMock()

				repository.PatchGetUser("test", entities.User{}, errors.New(database.ErrUserNotFound))

				return repository
			},
			loginUser: entities.UserLogin{
				Username: "test",
				Password: "test",
			},
			asserts: func(loginResponse *entities.LoginResponse, err error) {
				ass.NotNil(err)
				ass.EqualError(err, "invalid user")
				ass.Nil(loginResponse)
			},
		},
		{
			name: "Should Fail - Invalid Password",
			repository: func() userRepository.UserRepository {
				repository := mocks.NewMock()
				user := entities.User{
					ID:       uint(2),
					Username: "test",
					Password: "$2a$05$hoDkIPIevK0kIXmBK8ilEuz6iKifccSFy9ndir9CU.z/1ZIqy4ZCu",
				}
				repository.PatchGetUser("test", user, nil)

				return repository
			},
			loginUser: entities.UserLogin{
				Username: "test",
				Password: "invalid_password",
			},
			asserts: func(loginResponse *entities.LoginResponse, err error) {
				ass.Nil(loginResponse)
				ass.NotNil(err)
				ass.EqualError(err, "invalid user")
			},
		},
	}

	for _, test := range tests {
		// Given
		repository := test.repository()

		// When
		usecases := usecases.NewUseCasesMock(repository)
		loginResponse, err := usecases.LoginUser(test.loginUser)

		// Then
		test.asserts(loginResponse, err)
	}
}

func Test_AuthenticatedUser(t *testing.T) {
	ass := assert.New(t)

	tests := []struct {
		name       string
		repository func() userRepository.UserRepository
		token      func() string
		asserts    func(authenticatedResponse entities.AuthenticatedResponse, err error)
	}{
		{
			name: "Success",
			repository: func() userRepository.UserRepository {
				repository := mocks.NewMock()
				user := entities.User{
					ID:       uint(2),
					Username: "test",
					Password: "$2a$05$hoDkIPIevK0kIXmBK8ilEuz6iKifccSFy9ndir9CU.z/1ZIqy4ZCu",
				}
				repository.PatchGetUser("test", user, nil)

				return repository
			},
			token: func() string {
				repository := mocks.NewMock()
				user := entities.User{
					ID:       uint(2),
					Username: "test",
					Password: "$2a$05$hoDkIPIevK0kIXmBK8ilEuz6iKifccSFy9ndir9CU.z/1ZIqy4ZCu",
				}
				repository.PatchGetUser("test", user, nil)
				usecases := usecases.NewUseCasesMock(repository)
				loginUser := entities.UserLogin{
					Username: "test",
					Password: "test",
				}
				loginResponse, _ := usecases.LoginUser(loginUser)
				return loginResponse.Token
			},
			asserts: func(authenticatedResponse entities.AuthenticatedResponse, err error) {
				ass.True(authenticatedResponse.Authenticated)
				ass.Nil(err)
			},
		},
		{
			name: "Should Fail - Invalid Token",
			repository: func() userRepository.UserRepository {
				repository := mocks.NewMock()
				user := entities.User{
					ID:       uint(2),
					Username: "test",
					Password: "$2a$05$hoDkIPIevK0kIXmBK8ilEuz6iKifccSFy9ndir9CU.z/1ZIqy4ZCu",
				}
				repository.PatchGetUser("test", user, nil)

				return repository
			},
			token: func() string {
				return "invalid"
			},
			asserts: func(authenticatedResponse entities.AuthenticatedResponse, err error) {
				ass.False(authenticatedResponse.Authenticated)
				ass.NotNil(err)
			},
		},
	}

	for _, test := range tests {
		// Given
		repository := test.repository()

		// When
		usecases := usecases.NewUseCasesMock(repository)
		authenticatedResponse, err := usecases.AuthenticatedUser(test.token())

		// Then
		test.asserts(authenticatedResponse, err)
	}
}
