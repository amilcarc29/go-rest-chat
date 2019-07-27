package http

import (
	"encoding/json"
	"errors"
	"go-rest-chat/src/api/domain/user/entities"
	"strings"
)

// IsAuthenticated returns if the user is authenticated
func (repository *MessageHTTPRepository) IsAuthenticated(token string) (entities.AuthenticatedResponse, error) {

	var authenticated entities.AuthenticatedResponse

	if strings.Contains(token, "Bearer") || strings.Contains(token, "bearer") {
		tokenSplitted := strings.SplitAfter(token, " ")
		if len(tokenSplitted) < 2 {
			return entities.AuthenticatedResponse{}, errors.New("invalid token")
		}
		token = tokenSplitted[1]
	}

	resp, err := repository.client.R().
		SetAuthToken(token).
		Get("http://127.0.0.1:8080/authenticated")
	if err != nil {
		return authenticated, err
	}

	if resp.StatusCode() == 401 {
		return authenticated, errors.New("not authenticated")
	}

	json.Unmarshal(resp.Body(), &authenticated)
	return authenticated, nil
}
