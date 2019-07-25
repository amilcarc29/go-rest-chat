package http

import (
	"encoding/json"
	"fmt"
	"go-rest-chat/src/api/domain/message/entities"
	"net/http"
	"net/url"
	"strconv"
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

// GetResource handler for getting a resource
// func (handler *MessageHandler) GetResource(w http.ResponseWriter, r *http.Request) {

// 	id := mux.Vars(r)["id"]
// 	resource, err := handler.usecases.GetResource(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(&resource)
// 	w.WriteHeader(http.StatusOK)
// }

// CreateResource handler to create a resource
// func (handler *MessageHandler) CreateResource(w http.ResponseWriter, r *http.Request) {

// 	var resource entities.Resource
// 		w.WriteHeader(http.StatusBadRequest)
// 	json.NewDecoder(r.Body).Decode(&resource)
// 	newID, err := handler.usecases.CreateResource(resource)
// 	if err != nil {
// 		return
// 	}

// 	resourceIDResponse := struct {
// 		ID uint `json:"id"`
// 	}{
// 		ID: newID,
// 	}

// 	json.NewEncoder(w).Encode(&resourceIDResponse)
// 	w.WriteHeader(http.StatusCreated)
// }

// DeleteResource handler to delete a resource
// func (handler *MessageHandler) DeleteResource(w http.ResponseWriter, r *http.Request) {

// 		w.WriteHeader(http.StatusBadRequest)
// 	id := mux.Vars(r)["id"]
// 	err := handler.usecases.DeleteResource(id)
// 	if err != nil {
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }
