package http

import (
	"go-rest-chat/src/api/infraestructure/dependencies"

	"github.com/gorilla/mux"
)

// CheckHTTPRepository defines a Check http repository struct
type CheckHTTPRepository struct {
	http *mux.Router
}

// NewRepository returns a new Check http repository
func NewRepository(container *dependencies.Container) *CheckHTTPRepository {

	return &CheckHTTPRepository{
		http: container.RouterHandler(),
	}
}
