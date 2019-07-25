package entities

// Message defines a struct that represents a Message
type Message struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	SenderID    string  `json:"sender_id"`
	RecipientID string  `json:"recipient_id"`
	Content     Content `json:"content_id"`
	ContentID   uint    `json:"-"`
}

// Content defines a struct that represents a content in a message
type Content struct {
	ID   uint   `gorm:"primary_key" json:"-"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type Error struct {
	Error string `json:"error"`
}
