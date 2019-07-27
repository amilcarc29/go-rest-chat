package usecases_test

import (
	"encoding/json"
	"errors"
	"go-rest-chat/src/api/domain/message/entities"
	"go-rest-chat/src/api/domain/message/usecases"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_GetMessages_Success(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	// Mocks DB
	dateString := "2019-07-25T20:49:00Z"
	date, _ := time.Parse(time.RFC3339, dateString)
	senderID := uint(1)
	recipientID := uint(2)
	start := uint(2)
	limit := uint(2)
	rows := []map[string]interface{}{
		{"id": uint(2), "timestamp": date, "sender": senderID, "recipient": recipientID, "content": `{"type":"image","url":"imgurl.com","height":123,"width":123}`},
		{"id": uint(3), "timestamp": date, "sender": senderID, "recipient": recipientID, "content": `{"type": "text", "text": "text test 1"}`},
	}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "messages"`).WithReply(rows)

	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(200, `{"authenticated": true, "id": "1"}`))

	// Then
	messages, err := usecases.GetMessages("token", recipientID, start, limit)
	assert.NotNil(messages)
	assert.Len(messages, 2)
	assert.Nil(err)
}

func Test_GetMessages_NotAuthenticated_Fail(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	// Mocks DB
	dateString := "2019-07-25T20:49:00Z"
	date, _ := time.Parse(time.RFC3339, dateString)
	senderID := uint(1)
	recipientID := uint(2)
	start := uint(2)
	limit := uint(2)
	rows := []map[string]interface{}{
		{"id": uint(2), "timestamp": date, "sender": senderID, "recipient": recipientID, "content": `{"type":"image","url":"imgurl.com","height":123,"width":123}`},
		{"id": uint(3), "timestamp": date, "sender": senderID, "recipient": recipientID, "content": `{"type": "text", "text": "text test 1"}`},
	}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "messages"`).WithReply(rows)

	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(401, `{"authenticated": false, "id": "0"}`))

	// Then
	messages, err := usecases.GetMessages("token", recipientID, start, limit)
	assert.Nil(messages)
	assert.NotNil(err)
	assert.EqualError(err, "not authenticated")
}

func Test_GetMessages_ApplyingDefaultLimit_Success(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	// Mocks DB
	dateString := "2019-07-25T20:49:00Z"
	date, _ := time.Parse(time.RFC3339, dateString)
	senderID := uint(1)
	recipientID := uint(2)
	start := uint(2)
	limit := uint(0)
	rows := []map[string]interface{}{
		{"id": uint(2), "timestamp": date, "sender": senderID, "recipient": recipientID, "content": `{"type":"image","url":"imgurl.com","height":123,"width":123}`},
		{"id": uint(3), "timestamp": date, "sender": senderID, "recipient": recipientID, "content": `{"type": "text", "text": "text test 1"}`},
	}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "messages"`).WithReply(rows)

	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(200, `{"authenticated": true, "id": "1"}`))

	// Then
	messages, err := usecases.GetMessages("token", recipientID, start, limit)
	assert.NotNil(messages)
	assert.Len(messages, 2)
	assert.Nil(err)
}

func Test_GetMessages_ErrorGettingMessagesFromDB_Fail(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	// Mocks DB
	recipientID := uint(2)
	start := uint(2)
	limit := uint(2)
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "messages"`).WithError(errors.New("forced for test"))

	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(200, `{"authenticated": true, "id": "1"}`))

	// Then
	messages, err := usecases.GetMessages("token", recipientID, start, limit)
	assert.Nil(messages)
	assert.NotNil(err)
	assert.EqualError(err, "forced for test")
}

func Test_PostMessage_Success(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	// Mocks DB
	dateString := "2019-07-12T20:49:00Z"
	date, _ := time.Parse(time.RFC3339, dateString)
	senderID := uint(1)
	recipientID := uint(2)
	contentString := `{"type":"text","text":"test"}`
	var content entities.Content
	json.Unmarshal([]byte(contentString), &content)
	message := entities.Message{
		Sender:    senderID,
		Recipient: recipientID,
		Content:   content,
	}
	container.Catcher().Reset().NewMock().WithQuery(`INSERT INTO "messages" ("timestamp","sender","recipient","content") VALUES (?,?,?,?)`).WithArgs(date, senderID, recipientID, contentString)

	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(200, `{"authenticated": true, "id": "1"}`))

	// Then
	messageID, _, err := usecases.PostMessage("token", message)
	assert.Equal(uint(5577006791947779410), messageID)
	assert.Nil(err)
}

func Test_PostMessage_NotAuthenticated_Fail(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	// Mocks DB
	dateString := "2019-07-12T20:49:00Z"
	date, _ := time.Parse(time.RFC3339, dateString)
	senderID := uint(1)
	recipientID := uint(2)
	contentString := `{"type":"text","text":"test"}`
	var content entities.Content
	json.Unmarshal([]byte(contentString), &content)
	message := entities.Message{
		Sender:    senderID,
		Recipient: recipientID,
		Content:   content,
	}
	container.Catcher().Reset().NewMock().WithQuery(`INSERT INTO "messages" ("timestamp","sender","recipient","content") VALUES (?,?,?,?)`).WithArgs(date, senderID, recipientID, contentString)

	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(401, `{"authenticated": false, "id": "0"}`))

	// Then
	messageID, _, err := usecases.PostMessage("token", message)
	assert.Equal(uint(0), messageID)
	assert.NotNil(err)
	assert.EqualError(err, "not authenticated")
}

func Test_PostMessage_ErrorCreatingMessage_Fail(t *testing.T) {
	// Given
	assert, container, usecases := getDependenciesMock(t)
	defer closeDependencies(container)

	// When
	// Mocks DB
	senderID := uint(1)
	recipientID := uint(2)
	contentString := `{"type":"text","text":"test"}`
	var content entities.Content
	json.Unmarshal([]byte(contentString), &content)
	message := entities.Message{
		Sender:    senderID,
		Recipient: recipientID,
		Content:   content,
	}
	container.Catcher().Reset().NewMock().WithQuery(`INSERT`).WithError(errors.New("forced for test"))

	// Mocks Rest Client
	httpmock.ActivateNonDefault(container.HTTPClient().GetClient())
	httpmock.RegisterResponder("GET", "http://127.0.0.1:8080/authenticated", httpmock.NewStringResponder(200, `{"authenticated": true, "id": "1"}`))

	// Then
	_, _, err := usecases.PostMessage("token", message)
	assert.NotNil(err)
	assert.EqualError(err, "forced for test")
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
	httpmock.DeactivateAndReset()
}
