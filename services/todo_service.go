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
	DeleteTodoById(id int) (*help.JsonResponse, error)
	UpdateTodoByid(id int, title, description, file string) (*help.JsonResponse, error)
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
	if title == "" {
		return help.HandlerError(400, "Title is required", nil), nil
	}

	if description == "" {
		return help.HandlerError(400, "description is required", nil), nil
	}

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
func (t *todoService) DeleteTodoById(id int) (*help.JsonResponse, error) {
	row, err := t.todoRepository.Delete(id)

	if row == 0 {
		return help.HandlerError(400, "Bad Request", nil), nil
	}

	if err != nil {
		return help.HandlerError(500, "Server Error", nil), err
	}

	return help.HandlerSuccess(200, "Delete Todo Success", nil), nil

}
func (t *todoService) UpdateTodoByid(id int, title, description, file string) (*help.JsonResponse, error) {
	todo, _, err := t.todoRepository.FindOneTodoByid(id)

	if err != nil {
		return help.HandlerError(400, "Todo Not Found", nil), err
	}

	if title != "" {
		todo.Title = title
	}
	if description != "" {
		todo.Description = description
	}
	if file != "" {
		todo.File = file
	}

	newTodo, row, err := t.todoRepository.Update(todo)

	if row == 0 {
		return help.HandlerError(400, "Bad Request", nil), nil
	}

	if err != nil {
		return help.HandlerError(500, "Server Error", nil), err
	}

	return help.HandlerSuccess(200, "Get data success", newTodo), nil

}
