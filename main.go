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

	sr := repositories.NewSubtodoRepository(&repositories.SRConfig{
		DB: database.Get(),
	})

	ts := services.NewTodoService(&services.TSConfig{
		TodoRepository: tr,
	})

	ss := services.NewSubtodoService(&services.SSConfig{
		SubodoRepository: sr,
		Ts:               ts,
	})

	h := handlers.New(&handlers.HandlerConfig{
		TodoService:    ts,
		SubtodoService: ss,
	})

	e.GET("/", HelloWorld)
	e.POST("/todo", h.AddTodo)
	e.GET("/todo", h.GetAllTodo)
	e.GET("/todo/subtodo", h.GetAllTodoWithSubtodo)
	e.DELETE("/todo/:id", h.DeleteTodoById)

	e.POST("/subtodo/:id", h.AddSubtodo)
	e.GET("/subtodo/:id", h.GetAllSubTodoByTodoId)
	e.DELETE("/subtodo/:id", h.DeleteSubtodoById)
	e.Logger.Fatal(e.Start(":8080"))
}
