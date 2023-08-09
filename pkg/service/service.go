package service

import (
	"github.com/tursynkhan/todo-app"
	"github.com/tursynkhan/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User)(int,error)
	GenerateToken(username,password string)(string,error)
	ParseToken(token string)(int, error)
}
type TodoList interface {
	Create(userId int,list todo.Todolist)(int,error)
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