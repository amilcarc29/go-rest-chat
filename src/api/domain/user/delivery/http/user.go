package http

import (
	"encoding/json"
	"go-rest-chat/src/api/domain/user/entities"
	"net/http"
	"strings"
)

// CreateUser handler to register a new user
func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser entities.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}

	id, err := handler.usecases.CreateUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}

	newUserIDResponse := entities.NewUserIDResponse{
		ID: id,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&newUserIDResponse)
}

// LoginUser handler to login an user
func (handler *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	var userLogin entities.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}

	loginResponse, err := handler.usecases.LoginUser(userLogin)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginResponse)
}

// AuthenticatedUser handler to check if an user is authenticated
func (handler *UserHandler) AuthenticatedUser(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")
	if strings.Contains(tokenString, "Bearer") || strings.Contains(tokenString, "bearer") {
		tokenSplitted := strings.SplitAfter(tokenString, " ")
		if len(tokenSplitted) < 1 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(entities.Error{
				Error: "invalid token",
			})
			return
		}
		tokenString = tokenSplitted[1]
	}

	authenticated, err := handler.usecases.AuthenticatedUser(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(entities.Error{
			Error: err.Error(),
		})
		return
	}

	if authenticated.Authenticated {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&authenticated)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}
