package main

import (
	"github.com/labstack/echo/v4"

	"todo_list/handler"
	"todo_list/models"
)

func main() {
	data := []models.TodoItem{}

	app := echo.New()
	app.Static("/dist", "dist")

	todoHandler := handler.TodoHandler{
		Items: &data,
	}

	app.GET("/", todoHandler.RenderIndex)

	app.POST("/add", todoHandler.RenderAddTodo)

	app.DELETE("/delete/:id", todoHandler.RenderDeleteTodo)

	port := ":8080"

	app.Logger.Fatal(app.Start(port))
}
