package http

import (
	"encoding/json"
	"go-rest-chat/src/api/domain/user/entities"
)

// IsAuthenticated returns if the user is authenticated
func (repository *MessageHTTPRepository) IsAuthenticated(token string) (entities.AuthenticatedResponse, error) {

	var authenticated entities.AuthenticatedResponse

	resp, err := repository.client.R().
		SetHeader("Authorization", token).
		Get("http://127.0.0.1:8080/authenticated")
	if err != nil {
		return authenticated, err
	}

	if resp.StatusCode() == 401 {
		return authenticated, nil
	}

	json.Unmarshal(resp.Body(), &authenticated)
	return authenticated, nil
}
