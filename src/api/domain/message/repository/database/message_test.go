package database_test

import (
	"encoding/json"
	"errors"
	"go-rest-chat/src/api/domain/message/entities"
	"go-rest-chat/src/api/domain/message/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetResources_Success(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer container.Database().Close()

	// When
	dateString := "2019-07-25T20:49:00Z"
	date, _ := time.Parse(time.RFC3339, dateString)
	senderID := uint(1)
	recipientID := uint(2)
	start := uint(2)
	limit := uint(2)
	contentMessage1 := entities.Content{
		Type:   "image",
		URL:    "imgurl.com",
		Height: 123,
		Width:  123,
	}
	contentMessage2 := entities.Content{
		Type: "text",
		Text: "text test 1",
	}
	rows := []map[string]interface{}{
		{"id": uint(2), "timestamp": date, "sender": senderID, "recipient": recipientID, "content": contentMessage1},
		{"id": uint(3), "timestamp": date, "sender": senderID, "recipient": recipientID, "content": contentMessage2},
	}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "messages"`).WithReply(rows)

	// Then
	messages, err := repository.GetMessages(senderID, recipientID, start, limit)
	assert.Len(messages, 2)
	assert.Equal(uint(2), messages[0].ID)
	assert.Equal("image", messages[0].Content.Type)
	assert.Nil(err)
}

func Test_GetResources_NoResults_Success(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer container.Database().Close()

	// When
	senderID := uint(1)
	recipientID := uint(2)
	start := uint(4)
	limit := uint(2)
	rows := []map[string]interface{}{}
	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "messages"`).WithReply(rows)

	// Then
	messages, err := repository.GetMessages(senderID, recipientID, start, limit)
	assert.Len(messages, 0)
	assert.Nil(err)
}

func Test_GetResources_NoResults_Fail(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer container.Database().Close()

	// When
	senderID := uint(1)
	recipientID := uint(2)
	start := uint(4)
	limit := uint(2)

	container.Catcher().Reset().NewMock().WithQuery(`SELECT * FROM "messages"`).WithError(errors.New("forced for test"))

	// Then
	messages, err := repository.GetMessages(senderID, recipientID, start, limit)
	assert.Len(messages, 0)
	assert.NotNil(err)
	assert.EqualError(err, "forced for test")
}

func Test_PutMessage_Success(t *testing.T) {
	// Given
	assert, container, repository := getDependenciesMock(t)
	defer container.Database().Close()

	// When
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

	// Then
	messageID, _, err := repository.PutMessage(container.Clock().Now(), message)
	assert.Nil(err)
	assert.Equal(uint(5577006791947779410), messageID)
}

func getDependenciesMock(t *testing.T) (*assert.Assertions, *dependencies.Container, repository.MessageRepository) {
	assert := assert.New(t)
	mockContainer, err := dependencies.NewMockContainer()
	assert.Nil(err)
	repository := repository.NewMessageRepository(mockContainer)
	return assert, mockContainer, repository
}
