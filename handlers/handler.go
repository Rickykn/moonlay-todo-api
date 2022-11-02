package handlers

import "github.com/Rickykn/moonlay-todo-api/services"

type Handler struct {
	todoService    services.TodoService
	subtodoService services.SubtodoService
}

type HandlerConfig struct {
	TodoService    services.TodoService
	SubtodoService services.SubtodoService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		todoService:    c.TodoService,
		subtodoService: c.SubtodoService,
	}
}
