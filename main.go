package main

import (
	"log"
	"net/http"

	"github.com/Rickykn/moonlay-todo-api/database"
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
	e.GET("/", HelloWorld)
	e.Logger.Fatal(e.Start(":8080"))
}
