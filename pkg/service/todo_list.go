package service

import (
	"github.com/tursynkhan/todo-app"
	"github.com/tursynkhan/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.Todolist) (int, error) {
	return s.repo.Create(userId,list)
}

func (s *TodoListService)GetAll(userId int)([]todo.Todolist,error){
	return s.repo.GetAll(userId)
}
func (s *TodoListService)GetById(userId,listId int)(todo.Todolist,error){
	return s.repo.GetById(userId,listId)
}
func(s *TodoListService)Delete(userId,listId int)error{
	return s.repo.Delete(userId,listId)
}
func (s *TodoListService)Update(userId, listId int,input todo.UpdateLisInput)error{
	if err:=input.Validate();err!=nil{
		return err
	}
	return s.repo.Update(userId,listId,input)
}