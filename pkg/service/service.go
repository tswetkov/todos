package service

import (
	"github.com/tswetkov/todos"
	"github.com/tswetkov/todos/pkg/repository"
)

type Authorization interface {
	CreateUser(user todos.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
