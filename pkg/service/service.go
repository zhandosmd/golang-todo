package service

import (
	todo "github.com/zhandosmd/golang-todo"
	"github.com/zhandosmd/golang-todo/pkg/repository"
)

type Authorization interface {
	CreateUser(User todo.User) (int, error)
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
