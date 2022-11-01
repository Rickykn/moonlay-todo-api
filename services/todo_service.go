package services

import (
	help "github.com/Rickykn/moonlay-todo-api/helpers"
	r "github.com/Rickykn/moonlay-todo-api/repositories"
)

type TodoService interface {
	CreateTodo(title, description, file string) (*help.JsonResponse, error)
}

type todoService struct {
	todoRepository r.TodoRepository
}

type TSConfig struct {
	TodoRepository r.TodoRepository
}

func NewTodoService(c *TSConfig) TodoService {

	return &todoService{
		todoRepository: c.TodoRepository,
	}
}

func (t *todoService) CreateTodo(title, description, file string) (*help.JsonResponse, error) {
	todo, err := t.todoRepository.Create(title, description, file)

	if err != nil {
		return help.HandlerError(500, "Server Error", nil), err
	}

	return help.HandlerSuccess(201, "Success Add New Todo", todo), nil
}
