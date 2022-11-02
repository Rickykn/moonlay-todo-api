package handlers

import (
	"net/http"
	"strconv"

	"github.com/Rickykn/moonlay-todo-api/helpers"
	"github.com/labstack/echo"
)

func (h *Handler) AddSubtodo(c echo.Context) error {
	var isSuccess = true
	var fileName string
	title := c.FormValue("title")
	description := c.FormValue("description")
	id := c.Param("id")
	convId, _ := strconv.Atoi(id)
	file, err := c.FormFile("file")

	if file != nil {
		fileName, isSuccess = helpers.Uploader(file, err)
	}

	if isSuccess {
		response, _ := h.subtodoService.CreateSubTodo(title, description, fileName, convId)
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

func (h *Handler) GetAllSubTodoByTodoId(c echo.Context) error {
	id := c.Param("id")
	convId, _ := strconv.Atoi(id)

	response, _ := h.subtodoService.GetAllTodoByTodoId(convId)

	if response.Error {
		return c.JSON(response.Code, response)
	} else {
		return c.JSON(response.Code, response)
	}
}

func (h *Handler) DeleteSubtodoById(c echo.Context) error {
	id := c.Param("id")
	convId, _ := strconv.Atoi(id)

	response, _ := h.subtodoService.DeleteSubtodoById(convId)

	if response.Error {
		return c.JSON(response.Code, response)
	} else {
		return c.JSON(response.Code, response)
	}
}
