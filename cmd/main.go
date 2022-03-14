package main

import (
	"log"

	"github.com/tswetkov/todos"
)

func main() {
	server := new(todos.Server)

	if err := server.Run("8000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
