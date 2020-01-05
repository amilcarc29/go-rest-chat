package database

import (
	"go-rest-chat/src/api/domain/message/entities"
	"time"
)

const (
	errMessageNotFound = "Message %s not found"
)

// GetMessages returns all the messages found
func (repository *MessageDatabaseRepository) GetMessages(sender, recipient, start, limit uint) ([]entities.Message, error) {
	var messages []entities.Message
	if err := repository.database.Where("sender = ? AND recipient = ? AND id >= ?", sender, recipient, start).Limit(limit).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

// PutMessage creates a new message
func (repository *MessageDatabaseRepository) PutMessage(now time.Time, message entities.Message) (uint, time.Time, error) {
	message.Timestamp = now
	if err := repository.database.Create(&message).Error; err != nil {
		return 0, time.Time{}, err
	}
	return message.ID, message.Timestamp, nil
}
