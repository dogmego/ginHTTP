package repository

import (
	"github.com/jmoiron/sqlx"
	"maxHTTP"
)

type Authorization interface {
	CreateUser(user maxHTTP.User) (int, error)
	GetUser(username, password string) (maxHTTP.User, error)
}

type TodoList interface {
	Create(userId int, list maxHTTP.TodoList) (int, error)
	GetAll(userId int) ([]maxHTTP.TodoList, error)
	GetById(userId, listId int) (maxHTTP.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input maxHTTP.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item maxHTTP.TodoItem) (int, error)
	GetAll(userId, listId int) ([]maxHTTP.TodoItem, error)
	GetById(userId, itemId int) (maxHTTP.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input maxHTTP.UpdateItemInput) error
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
