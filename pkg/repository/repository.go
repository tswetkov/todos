package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tswetkov/todos"
)

type Authorization interface {
	CreateUser(user todos.User) (int, error)
	GetUser(username, password string) (todos.User, error)
}

type TodoList interface {
	Create(userId int, list todos.TodoList) (int, error)
	GetAll(userId int) ([]todos.TodoList, error)
	GetById(userId int, listId int) (todos.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, list todos.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todos.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todos.TodoItem, error)
	GetById(userId, itemId int) (todos.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todos.UpdateItemInput) error
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
