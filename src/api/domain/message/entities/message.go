package entities

// Message defines a struct that represents a Message
type Message struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	SenderID    string  `json:"sender_id"`
	RecipientID string  `json:"recipient_id"`
	Content     Content `json:"content"`
	ContentID   uint    `json:"-"`
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

// DBContent defines a struct that represents a content saved in the db
type DBContent struct {
	ID       uint   `gorm:"primary_key" json:"-"`
	Type     string `json:"type"`
	Metadata string `json:"metadata"`
}

// TableName overrides the table name for DBContent struct
func (DBContent) TableName() string {
	return "contents"
}

// Error defines a struct for error
type Error struct {
	Error string `json:"error"`
}
