package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tswetkov/todos"
)

type Authorization interface {
	CreateUser(user todos.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
