package database_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-rest-chat/src/api/domain/message/entities"
	"go-rest-chat/src/api/domain/message/repository"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
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
	rows := sqlmock.NewRows([]string{"id", "timestamp", "sender", "recipient", "content"}).
		AddRow(uint(2), date, senderID, recipientID, `{"type":"image","url":"imgurl.com","height":123,"width":123}`).
		AddRow(uint(3), date, senderID, recipientID, `{"type": "text","text": "text test 1"}`)

	queryQuoted := regexp.QuoteMeta(`SELECT * FROM "messages" WHERE (sender = ? AND recipient = ? AND id >= ?) LIMIT 2`)
	(*container.SQLMock()).ExpectQuery(queryQuoted).WithArgs(senderID, recipientID, start).WillReturnRows(rows)

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
	rows := sqlmock.NewRows([]string{"id", "timestamp", "sender", "recipient", "content"})

	queryQuoted := regexp.QuoteMeta(`SELECT * FROM "messages" WHERE (sender = ? AND recipient = ? AND id >= ?) LIMIT 2`)
	(*container.SQLMock()).ExpectQuery(queryQuoted).WithArgs(senderID, recipientID, start).WillReturnRows(rows)

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

	queryQuoted := regexp.QuoteMeta(`SELECT * FROM "messages" WHERE (sender = ? AND recipient = ? AND id >= ?) LIMIT 2`)
	(*container.SQLMock()).ExpectQuery(queryQuoted).WithArgs(senderID, recipientID, start).WillReturnError(errors.New("forced for test"))

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

	queryQuoted := regexp.QuoteMeta(`INSERT INTO "messages" ("timestamp","sender","recipient","content") VALUES (?,?,?,?) RETURNING "messages"."id"`)
	(*container.SQLMock()).ExpectBegin()
	(*container.SQLMock()).ExpectQuery(queryQuoted).WithArgs(date, senderID, recipientID, contentString).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Then
	messageID, _, err := repository.PutMessage(container.Clock().Now(), message)
	assert.Nil(err)
	fmt.Println(messageID)
}

func getDependenciesMock(t *testing.T) (*assert.Assertions, *dependencies.Container, repository.MessageRepository) {
	assert := assert.New(t)
	mockContainer, err := dependencies.NewMockContainer()
	assert.Nil(err)
	repository := repository.NewMessageRepository(mockContainer)
	return assert, mockContainer, repository
}
