package services

import (
	help "github.com/Rickykn/moonlay-todo-api/helpers"
	r "github.com/Rickykn/moonlay-todo-api/repositories"
)

type SubtodoService interface {
	CreateSubTodo(title, description, file string, id int) (*help.JsonResponse, error)
}

type subtodoService struct {
	subtodoRepository r.SubtodoRepository
	ts                TodoService
}

type SSConfig struct {
	SubodoRepository r.SubtodoRepository
	Ts               TodoService
}

func NewSubtodoService(c *SSConfig) SubtodoService {

	return &subtodoService{
		subtodoRepository: c.SubodoRepository,
		ts:                c.Ts,
	}
}

func (s *subtodoService) CreateSubTodo(title, description, file string, id int) (*help.JsonResponse, error) {
	todo, errTodo := s.ts.GetTodo(id)

	if errTodo != nil {
		return help.HandlerError(404, "Todo Not Found", nil), errTodo
	}

	subtodo, err := s.subtodoRepository.Create(title, description, file, id)
	if err != nil {
		return help.HandlerError(500, "Server Error", nil), err
	}
	subtodo.Todo = *todo

	return help.HandlerSuccess(201, "Success Add New SubTodo", subtodo), nil

}
