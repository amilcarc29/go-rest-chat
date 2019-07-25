package database

import (
	"encoding/json"
	"go-rest-chat/src/api/domain/message/entities"
)

const (
	errMessageNotFound = "Message %s not found"
)

// GetMessages returns all the messages found
func (repository *MessageDatabaseRepository) GetMessages(sender, recipient, start, limit uint) ([]entities.Message, error) {
	var messages []entities.Message
	if err := repository.database.Where("sender_id = ? AND recipient_id = ? AND id >= ?", sender, recipient, start).Limit(limit).Find(&messages).Error; err != nil {
		return nil, err
	}
	for i := range messages {
		var dbContent entities.DBContent
		if err := repository.database.Where("id = ?", messages[i].ContentID).First(&dbContent).Error; err != nil {
			return nil, err
		}
		messages[i].Content = getMessageContent(dbContent)
	}
	return messages, nil
}

func getMessageContent(dbContent entities.DBContent) entities.Content {
	switch dbContent.Type {
	case "text":
		var textContent entities.Content
		json.Unmarshal([]byte(dbContent.Metadata), &textContent)
		textContent.Type = dbContent.Type
		return textContent
	case "image":
		var imageContent entities.Content
		json.Unmarshal([]byte(dbContent.Metadata), &imageContent)
		imageContent.Type = dbContent.Type
		return imageContent
	case "video":
		var videoContent entities.Content
		json.Unmarshal([]byte(dbContent.Metadata), &videoContent)
		videoContent.Type = dbContent.Type
		return videoContent
	}
	return entities.Content{}
}

// GetResource returns the resource found
// func (repository *MessageDatabaseRepository) GetResource(id string) (entities.Message, error) {
// 	var message entities.Message
// 	idConverted, _ := strconv.Atoi(id)
// 	err := repository.database.First(&message, idConverted).Error
// 	if message.ID == 0 || err != nil {
// 		return message, fmt.Errorf(errResourceNotFound, id)
// 	}
// 	return message, nil
// }

// PutResource inserts a new resource
// func (repository *MessageDatabaseRepository) PutResource(resource entities.Message) (uint, error) {
// 	if err := repository.database.Create(&resource).Error; err != nil {
// 		return 0, err
// 	}
// 	return resource.ID, nil
// }

// DeleteResource deletes a resource
// func (repository *MessageDatabaseRepository) DeleteResource(id string) error {
// 	resource, _ := repository.GetResource(id)
// 	if err := repository.database.Delete(&resource).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
