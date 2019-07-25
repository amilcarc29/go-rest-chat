package http

import (
	"go-rest-chat/src/api/infraestructure/dependencies"

	"github.com/go-resty/resty"
)

// MessageHTTPRepository defines a message http repository struct
type MessageHTTPRepository struct {
	client *resty.Client
}

// NewRepository returns a new message http repository
func NewRepository(container *dependencies.Container) *MessageHTTPRepository {

	return &MessageHTTPRepository{
		client: container.HTTPClient(),
	}
}
