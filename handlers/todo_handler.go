package handlers

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Rickykn/moonlay-todo-api/helpers"
	"github.com/labstack/echo"
)

func (h *Handler) AddTodo(c echo.Context) error {

	isSuccess := true
	var fileType, fileName string
	title := c.FormValue("title")
	description := c.FormValue("description")
	file, err := c.FormFile("file")

	if err != nil {
		isSuccess = false
	} else {
		src, err := file.Open()
		if err != nil {
			isSuccess = false
		} else {
			fileByte, _ := ioutil.ReadAll(src)
			fileType = http.DetectContentType(fileByte)
			if fileType == "application/pdf" {
				fileName = "uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
			} else if fileType == "text/plain; charset=utf-8" {
				fileName = "uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".txt"
			} else {
				isSuccess = false
			}

			err = ioutil.WriteFile(fileName, fileByte, 0777)
			if err != nil {
				isSuccess = false
			}

		}
		defer src.Close()
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
