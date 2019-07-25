package main

import (
	"fmt"
	checkHandler "go-rest-chat/src/api/domain/check/delivery/http"
	messageHandler "go-rest-chat/src/api/domain/message/delivery/http"
	userHandler "go-rest-chat/src/api/domain/user/delivery/http"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"log"
	"net/http"
)

func main() {
	container, err := dependencies.NewContainer()

	if err != nil {
		fmt.Printf(err.Error())
	}

	userHandler := userHandler.NewUserHandler(container)
	checkHandler := checkHandler.NewCheckHandler(container)
	messageHandler := messageHandler.NewMessageHandler(container)

	routerHandler := container.RouterHandler()

	// Check
	routerHandler.HandleFunc("/check", checkHandler.Check).Methods("GET")

	// User
	routerHandler.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	routerHandler.HandleFunc("/login", userHandler.LoginUser).Methods("POST")
	routerHandler.HandleFunc("/authenticated", userHandler.AuthenticatedUser).Methods("GET")

	// Messages
	// routerHandler.HandleFunc("/messages", messageHandler.LoginUser).Methods("POST")
	routerHandler.HandleFunc("/messages", messageHandler.GetMessages).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", routerHandler))
}
