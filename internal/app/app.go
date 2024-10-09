package app

import (
	"fmt"

	"api-demo/internal/controller"
	"api-demo/internal/repository"
)

func Run() {
	fmt.Println("Hello, World!")

	// TODO: add ENV variables
	// TODO: create database connection and pass it to repository
	// TODO: implement datbase with GORM
	// TODO: create migration script

	// Create user repository
	userRepository := repository.NewUserRepository()

	// Create user controller
	_ = controller.NewUserController(userRepository)

	// TODO: create http server
	// TODO: add graceful shutdown
}
