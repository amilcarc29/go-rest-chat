package mocks

import (
	"crypto/md5"
	"encoding/json"
	"go-rest-chat/src/api/domain/message/entities"
	userEntities "go-rest-chat/src/api/domain/user/entities"
)

// Mock struct for Repository mock
type Mock struct {
	patchGetMessagesMap     map[hash][]outputForGetMessages
	patchPutMessageMap      map[hash][]outputForPutMessage
	patchIsAuthenticatedMap map[hash][]outputForIsAuthenticated
}

// NewMock Repository Mock
func NewMock() *Mock {
	patchGetMessagesMap := make(map[hash][]outputForGetMessages)
	patchPutMessageMap := make(map[hash][]outputForPutMessage)
	patchIsAuthenticatedMap := make(map[hash][]outputForIsAuthenticated)

	return &Mock{
		patchGetMessagesMap:     patchGetMessagesMap,
		patchPutMessageMap:      patchPutMessageMap,
		patchIsAuthenticatedMap: patchIsAuthenticatedMap,
	}
}

type hash [16]byte

func toHash(input interface{}) hash {
	jsonBytes, _ := json.Marshal(input)
	return md5.Sum(jsonBytes)
}

// PatchGetMessages patch for GetMessages function
func (mock *Mock) PatchGetMessages(sender, recipient, start, limit uint, messages []entities.Message, outputError error) {
	inputHash := toHash(getInputForGetMessages(sender, recipient, start, limit))
	output := getOutputForGetMessages(messages, outputError)

	if _, exists := mock.patchGetMessagesMap[inputHash]; !exists {
		arrOutputForGetMessages := make([]outputForGetMessages, 0)
		mock.patchGetMessagesMap[inputHash] = arrOutputForGetMessages
	}
	mock.patchGetMessagesMap[inputHash] = append(mock.patchGetMessagesMap[inputHash], output)
}

// GetMessages mock for GetMessages function
func (mock *Mock) GetMessages(sender, recipient, start, limit uint) ([]entities.Message, error) {
	inputHash := toHash(getInputForGetMessages(sender, recipient, start, limit))
	arrOutputForGet, exists := mock.patchGetMessagesMap[inputHash]
	if !exists || len(arrOutputForGet) == 0 {
		panic("Mock not available for GetMessages")
	}

	output := arrOutputForGet[0]
	arrOutputForGet = arrOutputForGet[1:]
	mock.patchGetMessagesMap[inputHash] = arrOutputForGet

	return output.messages, output.error
}

type inputForGetMessage struct {
	sender    uint
	recipient uint
	start     uint
	limit     uint
}

func getInputForGetMessages(sender, recipient, start, limit uint) inputForGetMessage {
	return inputForGetMessage{
		sender:    sender,
		recipient: recipient,
		start:     start,
		limit:     limit,
	}
}

type outputForGetMessages struct {
	messages []entities.Message
	error    error
}

func getOutputForGetMessages(messages []entities.Message, outputError error) outputForGetMessages {
	return outputForGetMessages{
		messages: messages,
		error:    outputError,
	}
}

// PatchPutMessage patch for PutMessage function
func (mock *Mock) PatchPutMessage(sender, recipient, start, limit uint, messages []entities.Message, outputError error) {
	inputHash := toHash(getInputForPutMessage(sender, recipient, start, limit))
	output := getOutputForPutMessage(messages, outputError)

	if _, exists := mock.patchPutMessageMap[inputHash]; !exists {
		arrOutputForPutMessage := make([]outputForPutMessage, 0)
		mock.patchPutMessageMap[inputHash] = arrOutputForPutMessage
	}
	mock.patchPutMessageMap[inputHash] = append(mock.patchPutMessageMap[inputHash], output)
}

// PutMessage mock for PutMessage function
func (mock *Mock) PutMessage(sender, recipient, start, limit uint) ([]entities.Message, error) {
	inputHash := toHash(getInputForPutMessage(sender, recipient, start, limit))
	arrOutputForPutMessage, exists := mock.patchPutMessageMap[inputHash]
	if !exists || len(arrOutputForPutMessage) == 0 {
		panic("Mock not available for PutMessage")
	}

	output := arrOutputForPutMessage[0]
	arrOutputForPutMessage = arrOutputForPutMessage[1:]
	mock.patchPutMessageMap[inputHash] = arrOutputForPutMessage

	return output.messages, output.error
}

type inputForPutMessage struct {
	sender    uint
	recipient uint
	start     uint
	limit     uint
}

func getInputForPutMessage(sender, recipient, start, limit uint) inputForPutMessage {
	return inputForPutMessage{
		sender:    sender,
		recipient: recipient,
		start:     start,
		limit:     limit,
	}
}

type outputForPutMessage struct {
	messages []entities.Message
	error    error
}

func getOutputForPutMessage(messages []entities.Message, outputError error) outputForPutMessage {
	return outputForPutMessage{
		messages: messages,
		error:    outputError,
	}
}

// PatchIsAuthenticated patch for IsAuthenticated function
func (mock *Mock) PatchIsAuthenticated(token string, authenticatedResponse userEntities.AuthenticatedResponse, outputError error) {
	inputHash := toHash(getInputForIsAuthenticated(token))
	output := getOutputForIsAuthenticated(authenticatedResponse, outputError)

	if _, exists := mock.patchIsAuthenticatedMap[inputHash]; !exists {
		arrOutputForIsAuthenticated := make([]outputForIsAuthenticated, 0)
		mock.patchIsAuthenticatedMap[inputHash] = arrOutputForIsAuthenticated
	}
	mock.patchIsAuthenticatedMap[inputHash] = append(mock.patchIsAuthenticatedMap[inputHash], output)
}

// IsAuthenticated mock for IsAuthenticated function
func (mock *Mock) IsAuthenticated(token string) (userEntities.AuthenticatedResponse, error) {
	inputHash := toHash(getInputForIsAuthenticated(token))
	arrOutputForIsAuthenticated, exists := mock.patchIsAuthenticatedMap[inputHash]
	if !exists || len(arrOutputForIsAuthenticated) == 0 {
		panic("Mock not available for IsAuthenticated")
	}

	output := arrOutputForIsAuthenticated[0]
	arrOutputForIsAuthenticated = arrOutputForIsAuthenticated[1:]
	mock.patchIsAuthenticatedMap[inputHash] = arrOutputForIsAuthenticated

	return output.authenticatedResponse, output.error
}

type inputForIsAuthenticated struct {
	token string
}

func getInputForIsAuthenticated(token string) inputForIsAuthenticated {
	return inputForIsAuthenticated{
		token: token,
	}
}

type outputForIsAuthenticated struct {
	authenticatedResponse userEntities.AuthenticatedResponse
	error                 error
}

func getOutputForIsAuthenticated(authenticatedResponse userEntities.AuthenticatedResponse, outputError error) outputForIsAuthenticated {
	return outputForIsAuthenticated{
		authenticatedResponse: authenticatedResponse,
		error:                 outputError,
	}
}
