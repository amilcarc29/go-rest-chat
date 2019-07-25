package usecases

import (
	"errors"
	"go-rest-chat/src/api/domain/message/entities"
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

	messages, err := usecases.messageRepository.GetMessages(authenticated.ID, recipient, start, limit)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// GetResource returns the resource
// func (usecases *UseCases) GetResource(id string) (*entities.Resource, error) {
// 	resource, err := usecases.resourceRepository.GetResource(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &resource, nil
// }

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
