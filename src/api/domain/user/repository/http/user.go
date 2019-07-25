package http

import (
	"go-rest-chat/src/api/infraestructure/dependencies"

	"github.com/gorilla/mux"
)

// UserHTTPRepository defines a user http repository struct
type UserHTTPRepository struct {
	http *mux.Router
}

// NewRepository returns a new user http repository
func NewRepository(container *dependencies.Container) *UserHTTPRepository {

	return &UserHTTPRepository{
		http: container.RouterHandler(),
	}
}
