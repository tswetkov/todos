package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/tswetkov/todos"
	"github.com/tswetkov/todos/pkg/handler"
	"github.com/tswetkov/todos/pkg/repository"
	"github.com/tswetkov/todos/pkg/service"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Initializing error!")
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(todos.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Main error: %s", err.Error())
	}
}
