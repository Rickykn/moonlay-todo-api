package handlers

import (
	"net/http"
	"strconv"

	"github.com/Rickykn/moonlay-todo-api/dtos"
	"github.com/Rickykn/moonlay-todo-api/helpers"
	"github.com/labstack/echo"
)

func (h *Handler) AddTodo(c echo.Context) error {

	var isSuccess = true
	var fileName string
	title := c.FormValue("title")
	description := c.FormValue("description")
	file, err := c.FormFile("file")

	if file != nil {
		fileName, isSuccess = helpers.Uploader(file, err)
	}

	if isSuccess {
		response, _ := h.todoService.CreateTodo(title, description, fileName)
		if response.Error {
			return c.JSON(response.Code, response)
		} else {
			return c.JSON(response.Code, response)
		}
	}

	return c.JSON(http.StatusBadRequest, helpers.JsonResponse{
		Code:    500,
		Message: "just can upload pdf and txt file",
		Data:    nil,
		Error:   false,
	})
}

func (h *Handler) GetAllTodo(c echo.Context) error {
	var todoResponse []*dtos.ResponeTodoDTO
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")

	if limit == "" {
		limit = "10"
	}

	convPage, _ := strconv.Atoi(page)
	convSize, _ := strconv.Atoi(limit)

	query := &helpers.Query{
		Page:  convPage,
		Limit: convSize,
	}

	todos, response := h.todoService.GetAllTodo(query)

	if response.Error {
		return c.JSON(response.Code, response)
	} else {
		for _, val := range todos {
			todo := &dtos.ResponeTodoDTO{
				Title:       val.Title,
				Description: val.Description,
				File:        val.File,
			}

			todoResponse = append(todoResponse, todo)
		}
		response.Data = todoResponse

		return c.JSON(response.Code, response)
	}

}

func (h *Handler) GetAllTodoWithSubtodo(c echo.Context) error {
	_, response := h.todoService.GetAllTodoWithSubtodo()

	if response.Error {
		return c.JSON(response.Code, response)
	} else {

		return c.JSON(response.Code, response)
	}

}

func (h *Handler) DeleteTodoById(c echo.Context) error {
	id := c.Param("id")
	convId, _ := strconv.Atoi(id)

	response, _ := h.todoService.DeleteTodoById(convId)

	if response.Error {
		return c.JSON(response.Code, response)
	} else {
		return c.JSON(response.Code, response)
	}
}
