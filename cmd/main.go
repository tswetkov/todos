package main

import (
	"log"

	"github.com/tswetkov/todos"
	"github.com/tswetkov/todos/pkg/handler"
	"github.com/tswetkov/todos/pkg/repository"
	"github.com/tswetkov/todos/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(todos.Server)

	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Main error: %s", err.Error())
	}
}
