package mocks

import (
	"crypto/md5"
	"encoding/json"
	entities "go-rest-chat/src/api/domain/user/entities"
)

// Mock struct for Repository mock
type Mock struct {
	patchGetUserMap    map[hash][]outputForGetUser
	patchCreateUserMap map[hash][]outputForCreateUser
}

// NewMock Repository Mock
func NewMock() *Mock {
	patchGetUserMap := make(map[hash][]outputForGetUser)
	patchCreateUserMap := make(map[hash][]outputForCreateUser)

	return &Mock{
		patchGetUserMap:    patchGetUserMap,
		patchCreateUserMap: patchCreateUserMap,
	}
}

type hash [16]byte

func toHash(input interface{}) hash {
	jsonBytes, _ := json.Marshal(input)
	return md5.Sum(jsonBytes)
}

// PatchGetUser patch for GetUser function
func (mock *Mock) PatchGetUser(username string, outputUser entities.User, outputError error) {
	inputHash := toHash(getInputForGetUser(username))
	output := getOutputForGetUser(outputUser, outputError)

	if _, exists := mock.patchGetUserMap[inputHash]; !exists {
		arrOutputForGetUserUser := make([]outputForGetUser, 0)
		mock.patchGetUserMap[inputHash] = arrOutputForGetUserUser
	}
	mock.patchGetUserMap[inputHash] = append(mock.patchGetUserMap[inputHash], output)
}

// GetUser mock for GetUser function
func (mock *Mock) GetUser(username string) (entities.User, error) {
	inputHash := toHash(getInputForGetUser(username))
	arrOutputForGetUser, exists := mock.patchGetUserMap[inputHash]
	if !exists || len(arrOutputForGetUser) == 0 {
		panic("Mock not available for GetUser")
	}

	output := arrOutputForGetUser[0]
	arrOutputForGetUser = arrOutputForGetUser[1:]
	mock.patchGetUserMap[inputHash] = arrOutputForGetUser

	return output.user, output.error
}

type inputForGetUser struct {
	username string
}

func getInputForGetUser(username string) inputForGetUser {
	return inputForGetUser{
		username: username,
	}
}

type outputForGetUser struct {
	user  entities.User
	error error
}

func getOutputForGetUser(user entities.User, outputError error) outputForGetUser {
	return outputForGetUser{
		user:  user,
		error: outputError,
	}
}

// PatchCreateUser patch for CreateUser function
func (mock *Mock) PatchCreateUser(newUser entities.User, outputUserID uint, outputError error) {
	inputHash := toHash(getInputForCreateUser(newUser))
	output := getOutputForCreateUser(outputUserID, outputError)

	if _, exists := mock.patchCreateUserMap[inputHash]; !exists {
		arrOutputForCreateUser := make([]outputForCreateUser, 0)
		mock.patchCreateUserMap[inputHash] = arrOutputForCreateUser
	}
	mock.patchCreateUserMap[inputHash] = append(mock.patchCreateUserMap[inputHash], output)
}

// CreateUser mock for CreateUser function
func (mock *Mock) CreateUser(newUser entities.User) (uint, error) {
	inputHash := toHash(getInputForCreateUser(newUser))
	arrOutputForCreateUser, exists := mock.patchCreateUserMap[inputHash]
	if !exists || len(arrOutputForCreateUser) == 0 {
		panic("Mock not available for CreateUser")
	}

	output := arrOutputForCreateUser[0]
	arrOutputForCreateUser = arrOutputForCreateUser[1:]
	mock.patchCreateUserMap[inputHash] = arrOutputForCreateUser

	return output.newUserID, output.error
}

type inputForCreateUser struct {
	newUser entities.User
}

func getInputForCreateUser(newUser entities.User) inputForCreateUser {
	return inputForCreateUser{
		newUser: newUser,
	}
}

type outputForCreateUser struct {
	newUserID uint
	error     error
}

func getOutputForCreateUser(newUserID uint, outputError error) outputForCreateUser {
	return outputForCreateUser{
		newUserID: newUserID,
		error:     outputError,
	}
}
