package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
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
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Initializing app error! %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Loading env errors! %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Initializing database error!")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(todos.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Main error: %s", err.Error())
	}
}
