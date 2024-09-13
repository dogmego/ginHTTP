package service

import (
	"maxHTTP"
	"maxHTTP/pkg/repository"
)

type Authorization interface {
	CreateUser(user maxHTTP.User) (int, error)
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
