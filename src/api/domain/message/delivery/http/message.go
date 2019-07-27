package http

import (
	"encoding/json"
	"fmt"
	"go-rest-chat/src/api/domain/message/entities"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// GetMessages handler for getting messages
func (handler *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	recipient, err := getValueFromParam(r.URL, "recipient", true)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}
	if recipient == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: "Recipient id is 0",
		})
		return
	}
	start, err := getValueFromParam(r.URL, "start", true)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}
	if start == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: "Start is 0",
		})
		return
	}

	limit, err := getValueFromParam(r.URL, "limit", false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}

	messages, err := handler.usecases.GetMessages(tokenString, recipient, start, limit)
	if err != nil && err.Error() == "not authenticated" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(entities.Error{
			Error: "not authenticated",
		})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}
	messagesOutput := struct {
		Messages []entities.Message `json:"messages"`
	}{
		Messages: messages,
	}
	json.NewEncoder(w).Encode(&messagesOutput)
	w.WriteHeader(http.StatusOK)
}

func getValueFromParam(url *url.URL, paramName string, required bool) (uint, error) {
	queryParams := url.Query()
	valueFromParams, ok := queryParams[paramName]
	if !ok && required {
		return 0, fmt.Errorf("Missing '%s' url param", paramName)
	}
	if !ok {
		return 0, nil
	}
	value, err := strconv.Atoi(valueFromParams[0])
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

// PostMessage handler for posting a message
func (handler *MessageHandler) PostMessage(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	var message entities.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}

	messageID, timestamp, err := handler.usecases.PostMessage(tokenString, message)
	if err != nil && err.Error() == "not authenticated" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(entities.Error{
			Error: "not authenticated",
		})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}
	newMessageOutput := struct {
		ID        uint      `json:"id"`
		Timestamp time.Time `json:"timestamp"`
	}{
		ID:        messageID,
		Timestamp: timestamp,
	}
	json.NewEncoder(w).Encode(&newMessageOutput)
	w.WriteHeader(http.StatusOK)
}
