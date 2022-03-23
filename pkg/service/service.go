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
	Create(userId int, list todos.TodoList) (int, error)
	GetAll(userId int) ([]todos.TodoList, error)
	GetById(userId int, listId int) (todos.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input todos.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item todos.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todos.TodoItem, error)
	GetById(userId, itemId int) (todos.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todos.UpdateItemInput) error
}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
