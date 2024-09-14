package service

import (
	"maxHTTP"
	"maxHTTP/pkg/repository"
)

type Authorization interface {
	CreateUser(user maxHTTP.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list maxHTTP.TodoList) (int, error)
	GetAll(userId int) ([]maxHTTP.TodoList, error)
	GetById(userId, listId int) (maxHTTP.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input maxHTTP.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item maxHTTP.TodoItem) (int, error)
	GetAll(userId, listId int) ([]maxHTTP.TodoItem, error)
	GetById(userId, itemId int) (maxHTTP.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input maxHTTP.UpdateItemInput) error
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
