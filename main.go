package main

import (
	"log"
	"net/http"

	"github.com/Rickykn/moonlay-todo-api/database"
	"github.com/Rickykn/moonlay-todo-api/handlers"
	"github.com/Rickykn/moonlay-todo-api/repositories"
	"github.com/Rickykn/moonlay-todo-api/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	errConnect := database.Connect()

	if errConnect != nil {
		panic(errConnect)
	}

	e := echo.New()

	tr := repositories.NewTodoRepository(&repositories.TRConfig{
		DB: database.Get(),
	})

	ts := services.NewTodoService(&services.TSConfig{
		TodoRepository: tr,
	})

	h := handlers.New(&handlers.HandlerConfig{
		TodoService: ts,
	})

	e.GET("/", HelloWorld)
	e.POST("/todo", h.AddTodo)
	e.Logger.Fatal(e.Start(":8080"))
}
