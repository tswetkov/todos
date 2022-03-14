package main

import (
	"log"

	"github.com/tswetkov/todos"
	"github.com/tswetkov/todos/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	server := new(todos.Server)

	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
