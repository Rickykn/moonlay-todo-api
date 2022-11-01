package handlers

import "github.com/Rickykn/moonlay-todo-api/services"

type Handler struct {
	todoService services.TodoService
}

type HandlerConfig struct {
	TodoService services.TodoService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		todoService: c.TodoService,
	}
}
