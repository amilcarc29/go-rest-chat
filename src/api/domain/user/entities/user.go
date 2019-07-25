package entities

// UserLogin struct for user login
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User struct for user
type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse struct for a login response
type LoginResponse struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
}

// NewUserIDResponse struct for a new user id response
type NewUserIDResponse struct {
	ID uint `json:"id"`
}

// AuthenticatedResponse struct for a authenticated response
type AuthenticatedResponse struct {
	Authenticated bool `json:"authenticated"`
	ID            uint `json:"id"`
}

type Error struct {
	Error string `json:"error"`
}
