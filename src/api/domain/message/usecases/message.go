package usecases

import (
	"errors"
	"go-rest-chat/src/api/domain/message/entities"
	"time"
)

// GetMessages returns the messages
func (usecases *UseCases) GetMessages(token string, recipient, start, limit uint) ([]entities.Message, error) {
	authenticated, err := usecases.messageRepository.IsAuthenticated(token)
	if err != nil {
		return nil, err
	}
	if !authenticated.Authenticated {
		return nil, errors.New("not authenticated")
	}

	if limit == 0 {
		limit = 100
	}

	messages, err := usecases.messageRepository.GetMessages(authenticated.ID, recipient, start, limit)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// PostMessage posts a new message
func (usecases *UseCases) PostMessage(token string, message entities.Message) (uint, time.Time, error) {
	authenticated, err := usecases.messageRepository.IsAuthenticated(token)
	if err != nil {
		return 0, time.Time{}, err
	}
	if !authenticated.Authenticated {
		return 0, time.Time{}, errors.New("not authenticated")
	}
	messageID, timestamp, err := usecases.messageRepository.PutMessage(usecases.clock.Now(), message)
	if err != nil {
		return 0, time.Time{}, err
	}
	return messageID, timestamp, nil
}
