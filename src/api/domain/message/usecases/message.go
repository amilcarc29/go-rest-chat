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
// GetResource returns the resource
// func (usecases *UseCases) GetResource(id string) (*entities.Resource, error) {
// 	resource, err := usecases.resourceRepository.GetResource(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &resource, nil
// }

	messageID, timestamp, err := usecases.messageRepository.PutMessage(message)
	if err != nil {
		return 0, time.Time{}, err
	}
	return messageID, timestamp, nil
}
// CreateResource creates a new resource
// func (usecases *UseCases) CreateResource(resource entities.Resource) (uint, error) {
// 	id, err := usecases.resourceRepository.PutResource(resource)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }

// DeleteResource deletes the resource
// func (usecases *UseCases) DeleteResource(id string) error {
// 	err := usecases.resourceRepository.DeleteResource(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
