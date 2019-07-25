package entities

// Message defines a struct that represents a Message
type Message struct {
	ID            uint    `gorm:"primary_key" json:"id"`
	SenderID      string  `json:"sender_id"`
	RecipientID   string  `json:"recipient_id"`
	Content       Content `json:"content"`
	ContentString string  `gorm:"column:content" json:"-"`
}

// Content defines a struct for the content of a message
type Content struct {
	Type   string `json:"type"`
	Text   string `json:"text,omitempty"`
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
	Source string `json:"source,omitempty"`
}

// Error defines a struct for error
type Error struct {
	Error string `json:"error"`
}
