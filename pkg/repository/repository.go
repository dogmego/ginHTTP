package repository

import (
	"github.com/jmoiron/sqlx"
	"maxHTTP"
)

type Authorization interface {
	CreateUser(user maxHTTP.User) (int, error)
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
