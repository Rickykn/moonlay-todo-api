package services

import (
	help "github.com/Rickykn/moonlay-todo-api/helpers"
	"github.com/Rickykn/moonlay-todo-api/models"
	r "github.com/Rickykn/moonlay-todo-api/repositories"
)

type TodoService interface {
	CreateTodo(title, description, file string) (*help.JsonResponse, error)
	GetAllTodo(query *help.Query) ([]*models.Todo, *help.JsonResponse)
	GetTodo(id int) (*models.Todo, error)
	GetAllTodoWithSubtodo() ([]*models.Todo, *help.JsonResponse)
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

func (t *todoService) GetAllTodo(query *help.Query) ([]*models.Todo, *help.JsonResponse) {
	todos, err := t.todoRepository.FindAllTodo(query)

	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil)
	}

	return todos, help.HandlerSuccess(200, "Get data success", todos)
}

func (t *todoService) GetTodo(id int) (*models.Todo, error) {
	todo, _, err := t.todoRepository.FindOneTodoByid(id)

	return todo, err
}

func (t *todoService) GetAllTodoWithSubtodo() ([]*models.Todo, *help.JsonResponse) {
	todos, err := t.todoRepository.FindAllTodoWithSubTodo()

	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil)
	}

	return todos, help.HandlerSuccess(200, "Get data success", todos)
}
