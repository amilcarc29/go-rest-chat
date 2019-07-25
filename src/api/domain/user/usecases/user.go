package usecases

import (
	"errors"
	"fmt"
	"go-rest-chat/src/api/domain/user/entities"
	"go-rest-chat/src/api/domain/user/repository/database"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user
func (usecases *UseCases) CreateUser(newUser entities.User) (uint, error) {

	_, err := usecases.userRepository.GetUser(newUser.Username)
	if err != nil {
		// user not found => creates a new user
		if err.Error() == database.ErrUserNotFound {
			hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 5)
			if err != nil {
				return 0, err
			}

			newUser.Password = string(hash)
			id, err := usecases.userRepository.CreateUser(newUser)
			if err != nil {
				return 0, err
			}

			return id, nil
		}
		// an actual error happened
		return 0, err
	}
	// User already exists
	return 0, errors.New("Username already exists")
}

// LoginUser login
func (usecases *UseCases) LoginUser(userLogin entities.UserLogin) (*entities.LoginResponse, error) {

	var user entities.User
	user, err := usecases.userRepository.GetUser(userLogin.Username)
	if err != nil {
		return nil, errors.New("invalid user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return nil, errors.New("invalid user")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &entities.LoginResponse{
		ID:    user.ID,
		Token: tokenString,
	}, nil
}

// AuthenticatedUser profile
func (usecases *UseCases) AuthenticatedUser(tokenString string) (entities.AuthenticatedResponse, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return entities.AuthenticatedResponse{
			Authenticated: false,
			ID:            0,
		}, err
	}

	var user entities.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Username = claims["username"].(string)
		user.ID = uint(claims["id"].(float64))

		return entities.AuthenticatedResponse{
			Authenticated: true,
			ID:            user.ID,
		}, nil
	}

	return entities.AuthenticatedResponse{
		Authenticated: false,
		ID:            0,
	}, nil
}
